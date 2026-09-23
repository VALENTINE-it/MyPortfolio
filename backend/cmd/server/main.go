package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

type HealthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// corsAndSecurityMiddleware handles restricted CORS and attaches recommended security headers
func corsAndSecurityMiddleware(allowedOrigins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			for _, allowed := range allowedOrigins {
				if allowed == "*" || allowed == origin || strings.TrimRight(allowed, "/") == strings.TrimRight(origin, "/") {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
					w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
					w.Header().Set("Access-Control-Max-Age", "86400")
					break
				}
			}
		}

		// Security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// healthHandler handles GET /api/health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Method not allowed",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HealthResponse{
		Success: true,
		Message: "API is running",
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "./portfolio.db"
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	allowedOrigins := []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://localhost:3000",
	}
	if frontendURL != "" && frontendURL != "http://localhost:5173" {
		allowedOrigins = append(allowedOrigins, frontendURL)
	}

	// Initialize Database (SQLite server-side persistence)
	db, err := database.InitDB(dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	log.Printf("Database initialized successfully at %s", dbURL)

	// Initialize Repositories
	projectRepo := repositories.NewProjectRepository(db)
	skillRepo := repositories.NewSkillRepository(db)
	contactRepo := repositories.NewContactRepository(db)

	// Initialize Services
	projectService := services.NewProjectService(projectRepo)
	skillService := services.NewSkillService(skillRepo)
	contactService := services.NewContactService(contactRepo)

	// Initialize Handlers
	projectHandler := handlers.NewProjectHandler(projectService)
	skillHandler := handlers.NewSkillHandler(skillService)
	contactHandler := handlers.NewContactHandler(contactService)

	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/api/health", healthHandler)

	// Projects
	mux.HandleFunc("/api/projects", projectHandler.HandleProjects)
	mux.HandleFunc("/api/projects/", projectHandler.HandleProjects)

	// Skills
	mux.HandleFunc("/api/skills", skillHandler.HandleSkills)

	// Contact
	mux.HandleFunc("/api/contact", contactHandler.HandleContact)

	handler := corsAndSecurityMiddleware(allowedOrigins, mux)

	// Server with configured timeouts
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for errors coming from listener
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server listening on port %s (http://localhost:%s/api)...", port, port)
		serverErrors <- server.ListenAndServe()
	}()

	// Channel to listen for interrupt/terminate signal from OS (Graceful shutdown)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server startup failed: %v", err)
		}

	case sig := <-shutdown:
		log.Printf("Shutdown signal %v received: initiating graceful shutdown...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown failed: %v, forcing server to close", err)
			_ = server.Close()
		}
		log.Println("Server gracefully stopped.")
	}
}
