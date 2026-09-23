package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/handlers"
	"portfolio-backend/internal/repositories"
	"portfolio-backend/internal/services"
)

type HealthResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// corsAndSecurityMiddleware handles CORS and attaches recommended security headers
func corsAndSecurityMiddleware(allowedOrigin string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if allowedOrigin == "*" || allowedOrigin == origin || allowedOrigin == "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			}
		} else if allowedOrigin != "" {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
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
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	// Initialize Database
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

	handler := corsAndSecurityMiddleware(frontendURL, mux)

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on port %s (http://localhost:%s/api)...", port, port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
