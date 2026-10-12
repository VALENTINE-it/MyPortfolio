package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"portfolio-backend/internal/cache"
	"portfolio-backend/internal/config"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/monitoring"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
	"portfolio-backend/internal/tasks"
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Configuration error: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize structured logging
	logger := monitoring.InitLogger(cfg.Environment, os.Stdout)
	logger.Info("Starting portfolio backend server",
		"environment", cfg.Environment,
		"port", cfg.Port,
	)

	// 3. Initialize SQLite database, run migrations, and seed data
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		logger.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := database.CheckIntegrity(db); err != nil {
		logger.Warn("Database integrity check warning", "warning", err)
	} else {
		logger.Info("Database integrity check passed (healthy)")
	}

	// 4. In-Memory Cache
	memCache := cache.NewMemoryCache(1000, 10*time.Minute)
	defer memCache.Close()

	// 5. Initialize Repositories
	projectRepo := repositories.NewProjectRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	contactRepo := repositories.NewContactRepository(db)
	adminRepo := repositories.NewAdminRepository(db)
	sessionRepo := repositories.NewSessionRepository(db)

	// 6. Initialize Services
	projectService := services.NewProjectService(projectRepo)
	skillService := services.NewSkillService(skillRepo)
	contactService := services.NewContactService(contactRepo)
	authService := services.NewAuthService(adminRepo, sessionRepo, cfg.SessionTTL)

	// 7. Background Worker
	worker := tasks.NewBackgroundWorker(db, authService, "./backups")
	worker.Start()
	defer worker.Stop()

	// 8. Initialize Handlers
	healthHandler := monitoring.NewHealthHandler(db)
	authHandler := handlers.NewAuthHandler(authService, cfg)
	projectHandler := handlers.NewProjectHandler(projectService, memCache)
	skillHandler := handlers.NewSkillHandler(skillService, memCache)
	contactHandler := handlers.NewContactHandler(contactService)
	dashboardHandler := handlers.NewAdminDashboardHandler(db)

	// 9. Rate Limiters
	publicLimiter := middleware.NewIPRateLimiter(cfg.PublicRPM, cfg.PublicRPM/2, cfg)
	defer publicLimiter.Close()

	contactLimiter := middleware.NewIPRateLimiter(cfg.ContactRPM, 3, cfg)
	defer contactLimiter.Close()

	loginLimiter := middleware.NewIPRateLimiter(cfg.LoginRPM, 5, cfg)
	defer loginLimiter.Close()

	adminLimiter := middleware.NewIPRateLimiter(cfg.AdminRPM, 20, cfg)
	defer adminLimiter.Close()

	// 10. HTTP ServeMux Routing
	mux := http.NewServeMux()

	// System & Health Endpoints
	mux.HandleFunc("/health/live", healthHandler.LivenessHandler)
	mux.HandleFunc("/health/ready", healthHandler.ReadinessHandler)
	mux.HandleFunc("/api/health", healthHandler.LegacyHealthHandler)

	if cfg.MetricsEnabled {
		mux.HandleFunc("/metrics", monitoring.GetMetrics().PrometheusHandler())
	}

	// Public Portfolio Routes
	mux.Handle("/api/projects", publicLimiter.Middleware(http.HandlerFunc(projectHandler.HandleProjects)))
	mux.Handle("/api/projects/", publicLimiter.Middleware(http.HandlerFunc(projectHandler.HandleProjects)))

	mux.Handle("/api/skills", publicLimiter.Middleware(http.HandlerFunc(skillHandler.HandleSkills)))
	mux.Handle("/api/skills/", publicLimiter.Middleware(http.HandlerFunc(skillHandler.HandleSkills)))

	mux.Handle("/api/contact", contactLimiter.Middleware(http.HandlerFunc(contactHandler.HandleContact)))

	// Admin Authentication Routes
	mux.Handle("/api/admin/login", loginLimiter.Middleware(http.HandlerFunc(authHandler.HandleLogin)))
	mux.HandleFunc("/api/admin/logout", authHandler.HandleLogout)
	mux.Handle("/api/admin/me", middleware.RequireAdmin(http.HandlerFunc(authHandler.HandleMe)))
	mux.HandleFunc("/api/admin/csrf", authHandler.HandleCSRF)

	// Admin Protected CRUD Routes
	mux.Handle("/api/admin/dashboard", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(dashboardHandler.HandleDashboard))))

	mux.Handle("/api/admin/projects", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(projectHandler.HandleAdminProjects))))
	mux.Handle("/api/admin/projects/", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(projectHandler.HandleAdminProjects))))

	mux.Handle("/api/admin/skills", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(skillHandler.HandleAdminSkills))))
	mux.Handle("/api/admin/skills/", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(skillHandler.HandleAdminSkills))))

	mux.Handle("/api/admin/messages", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(contactHandler.HandleAdminMessages))))
	mux.Handle("/api/admin/messages/", middleware.RequireAdmin(adminLimiter.Middleware(http.HandlerFunc(contactHandler.HandleAdminMessages))))

	// 11. Compose Global Middleware Chain
	// Order: Recovery -> RequestID -> Logging -> SecurityHeaders -> CORS -> BodyLimit -> Authenticate -> CSRF -> Router
	var rootHandler http.Handler = mux
	rootHandler = middleware.CSRFProtection(cfg)(rootHandler)
	rootHandler = middleware.Authenticate(authService)(rootHandler)
	rootHandler = middleware.BodyLimit(cfg.MaxRequestBodyBytes)(rootHandler)
	rootHandler = middleware.CORS(cfg.AllowedOrigins, rootHandler)
	rootHandler = middleware.SecurityHeaders(cfg.IsProduction(), rootHandler)
	rootHandler = middleware.Logging(rootHandler)
	rootHandler = middleware.RequestID(rootHandler)
	rootHandler = middleware.Recovery(rootHandler)

	// 12. HTTP Server with Strict Production Timeouts
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.Port),
		Handler:           rootHandler,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	serverErrors := make(chan error, 1)

	go func() {
		logger.Info("HTTP server running and listening",
			"addr", server.Addr,
			"url", fmt.Sprintf("http://localhost:%s", cfg.Port),
		)
		serverErrors <- server.ListenAndServe()
	}()

	// 13. Graceful Shutdown on SIGINT / SIGTERM
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server terminated unexpectedly", "error", err)
			os.Exit(1)
		}

	case sig := <-shutdown:
		logger.Info("Shutdown signal received: starting graceful shutdown", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Graceful shutdown failed: forcing close", "error", err)
			_ = server.Close()
		}

		logger.Info("Server stopped cleanly")
	}
}
