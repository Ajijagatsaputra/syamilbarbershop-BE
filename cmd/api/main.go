package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"syamilbarbershop-be/internal/config"
	"syamilbarbershop-be/internal/database"
	"syamilbarbershop-be/internal/routes"
)

func main() {
	config.LoadConfig()

	// Parse flags
	migrateFlag := flag.Bool("migrate", false, "run database migration and exit")
	seedFlag := flag.Bool("seed", false, "run database seeding and exit")
	flag.Parse()

	// Connect to database
	database.ConnectDB()

	// Always run database migrations on startup to keep schema in sync
	database.MigrateDB()

	// Handle migration flag
	if *migrateFlag {
		return
	}

	// Handle seed flag
	if *seedFlag {
		log.Println("Starting database seeding...")
		database.SeedData()
		log.Println("Database seeding finished successfully.")
		return
	}

	// Setup router
	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create http.Server for graceful shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Setup channel to listen for interrupt signals to shut down gracefully
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Run server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %s\n", err)
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()

	// Notify of shutdown
	stop()
	log.Println("Shutting down server gracefully...")

	// 5-second timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server exited cleanly")
}
