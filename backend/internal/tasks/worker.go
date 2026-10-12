package tasks

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"
	"sync"
	"time"

	"portfolio-backend/internal/database"
	"portfolio-backend/internal/services"
)

// BackgroundWorker runs scheduled maintenance and background tasks
type BackgroundWorker struct {
	db          *sql.DB
	authService *services.AuthService
	backupDir   string
	stopCh      chan struct{}
	wg          sync.WaitGroup
}

func NewBackgroundWorker(db *sql.DB, authService *services.AuthService, backupDir string) *BackgroundWorker {
	if backupDir == "" {
		backupDir = "./backups"
	}
	return &BackgroundWorker{
		db:          db,
		authService: authService,
		backupDir:   backupDir,
		stopCh:      make(chan struct{}),
	}
}

// Start launches the background task loops
func (w *BackgroundWorker) Start() {
	w.wg.Add(2)

	// 1. Session cleanup task (runs every hour)
	go func() {
		defer w.wg.Done()
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-w.stopCh:
				return
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				deleted, err := w.authService.CleanupExpiredSessions(ctx)
				cancel()
				if err != nil {
					slog.Error("Background task: error cleaning expired sessions", "error", err)
				} else if deleted > 0 {
					slog.Info("Background task: cleaned expired sessions", "count", deleted)
				}
			}
		}
	}()

	// 2. Database maintenance & backup task (runs every 24 hours)
	go func() {
		defer w.wg.Done()
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-w.stopCh:
				return
			case <-ticker.C:
				w.RunMaintenance()
			}
		}
	}()

	slog.Info("Background worker started successfully")
}

// RunMaintenance runs database integrity check and an atomic backup snapshot
func (w *BackgroundWorker) RunMaintenance() {
	slog.Info("Running scheduled database maintenance...")

	if err := database.CheckIntegrity(w.db); err != nil {
		slog.Error("Database integrity check failed during maintenance", "error", err)
		return
	}

	backupFile := filepath.Join(w.backupDir, fmt.Sprintf("portfolio-backup-%s.db", time.Now().Format("20060102-150405")))
	if err := database.Backup(w.db, backupFile); err != nil {
		slog.Error("Scheduled database backup failed", "error", err)
		return
	}

	slog.Info("Scheduled database maintenance completed", "backup_file", backupFile)
}

// Stop signals all background tasks to finish and waits for graceful completion
func (w *BackgroundWorker) Stop() {
	close(w.stopCh)
	w.wg.Wait()
	slog.Info("Background worker stopped gracefully")
}
