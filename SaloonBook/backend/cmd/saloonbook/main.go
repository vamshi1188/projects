package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"saloonbook/internal/config"
	"saloonbook/internal/db"
	"saloonbook/internal/handlers"
	"saloonbook/internal/middleware"
	"saloonbook/internal/repository"
	"saloonbook/pkg/logger"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load()

	// Initialize logger
	logger.Init()

	// Load configuration
	cfg := config.Load()

	// Database connection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db.Connect(ctx)
	defer func() {
		if db.Conn != nil {
			db.Conn.Close(context.Background())
		}
	}()

	// Initialize repositories
	serviceRepo := repository.NewServiceRepository()
	bookingRepo := repository.NewBookingRepository()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(db.Conn)
	serviceHandler := handlers.NewServiceHandler(serviceRepo)
	bookingHandler := handlers.NewBookingHandler(bookingRepo)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CORS())

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler.Check)
		r.Get("/services", serviceHandler.List)
		r.Get("/bookings", bookingHandler.List)
		r.Post("/bookings", bookingHandler.Create)
	})

	// Optional static file serving for built frontend
	if cfg.ServeFrontend {
		if stat, err := os.Stat(cfg.FrontendDist); err == nil && stat.IsDir() {
			fileServer := http.FileServer(http.Dir(cfg.FrontendDist))
			r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				// Serve index.html for SPA routes (no file extension)
				if filepath.Ext(req.URL.Path) == "" && req.URL.Path != "/" {
					http.ServeFile(w, req, filepath.Join(cfg.FrontendDist, "index.html"))
					return
				}
				fileServer.ServeHTTP(w, req)
			}))
			logger.Info.Printf("Serving frontend from %s", cfg.FrontendDist)
		} else {
			logger.Error.Printf("Frontend directory not found: %s", cfg.FrontendDist)
		}
	}

	// HTTP server with timeouts
	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info.Printf("Server starting on port %s (env: %s)", cfg.Port, cfg.Environment)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("Shutdown signal received, gracefully stopping...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error.Printf("Server forced to shutdown: %v", err)
	}

	logger.Info.Println("Server stopped cleanly")
}
