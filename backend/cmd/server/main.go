package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"portfolio-backend/internal/config"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/middleware"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

func main() {
	// 1. Load and validate configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configuration error: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize structured logging with slog
	var logHandler slog.Handler
	if cfg.IsProduction() {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	slog.Info("Starting portfolio backend server...",
		"env", cfg.AppEnv,
		"port", cfg.Port,
		"db_path", cfg.DatabasePath,
		"trust_proxy", cfg.TrustProxy,
	)

	// 3. Initialize SQLite Database
	db, err := database.InitDB(cfg.DatabasePath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Database initialized successfully with WAL mode & connection pooling")

	// 4. Initialize Repositories
	projectRepo := repositories.NewProjectRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	contactRepo := repositories.NewContactRepository(db)

	// 5. Initialize Services
	projectService := services.NewProjectService(projectRepo)
	skillService := services.NewSkillService(skillRepo)
	contactService := services.NewContactService(contactRepo)

	// 6. Initialize Handlers
	projectHandler := handlers.NewProjectHandler(projectService)
	skillHandler := handlers.NewSkillHandler(skillService)
	contactHandler := handlers.NewContactHandler(contactService)

	// 7. Initialize Rate Limiters
	generalLimiter := middleware.NewMemoryRateLimiter(5 * time.Minute)
	defer generalLimiter.Close()

	contactLimiter := middleware.NewMemoryRateLimiter(5 * time.Minute)
	defer contactLimiter.Close()

	// 8. Register Routes using Go standard library net/http ServeMux
	mux := http.NewServeMux()

	// Health Check
	mux.HandleFunc("/api/health", handlers.HandleHealth)

	// Projects
	mux.HandleFunc("/api/projects", projectHandler.HandleProjects)
	mux.HandleFunc("/api/projects/", projectHandler.HandleProjects)

	// Skills
	mux.HandleFunc("/api/skills", skillHandler.HandleSkills)

	// Contact (Wrapped with dedicated sensitive rate limiter and body size limit)
	contactLimiterMiddleware := middleware.RateLimit(
		contactLimiter,
		cfg.RateLimitContactReq,
		cfg.RateLimitContactWindow,
		cfg.TrustProxy,
	)
	bodyLimitMiddleware := middleware.BodyLimit(cfg.MaxRequestBodyBytes)
	mux.Handle("/api/contact", bodyLimitMiddleware(contactLimiterMiddleware(http.HandlerFunc(contactHandler.HandleContact))))

	// 9. Build Global Middleware Chain:
	// PanicRecovery -> RequestID -> RequestLogger -> SecurityHeaders -> CORS -> GeneralRateLimit -> ServeMux
	generalRateLimitMiddleware := middleware.RateLimit(
		generalLimiter,
		cfg.RateLimitGeneralReq,
		cfg.RateLimitGeneralWindow,
		cfg.TrustProxy,
	)

	var handler http.Handler = mux
	handler = generalRateLimitMiddleware(handler)
	handler = middleware.CORS(cfg.AllowedOrigins, handler)
	handler = middleware.SecurityHeaders(cfg.IsProduction(), handler)
	handler = middleware.RequestLogger(cfg.TrustProxy, handler)
	handler = middleware.RequestIDMiddleware(handler)
	handler = middleware.PanicRecovery(handler)

	// 10. Configure HTTP Server with strict timeouts and header limits
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", cfg.Port),
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
	}

	// Channel to listen for listener errors
	serverErrors := make(chan error, 1)

	go func() {
		slog.Info(fmt.Sprintf("Server listening at http://localhost:%s", cfg.Port))
		serverErrors <- server.ListenAndServe()
	}()

	// 11. Graceful Shutdown listener (SIGINT, SIGTERM)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server listener failed", "error", err)
			os.Exit(1)
		}

	case sig := <-shutdown:
		slog.Info("Shutdown signal received: starting graceful shutdown...", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error("Graceful shutdown failed, forcing server close", "error", err)
			_ = server.Close()
		}
		slog.Info("Server stopped cleanly.")
	}
}
