package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"youtube-backend/api/routes"
	"youtube-backend/configs"
	"youtube-backend/pkg/db"
)

func main() {
	// Print current directory for debugging
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory: ", err)
	}
	log.Println("Current working directory:", wd)

	// Load configuration
	configs.LoadConfig()

	// Initialize MongoDB connection
	mongoConfig := db.MongoConfig{
		URI:  configs.MongoDBURI,
		Name: configs.MongoDBName,
	}

	database, err := db.InitMongoDB(mongoConfig)
	if err != nil {
		log.Fatal("Failed to initialize MongoDB: ", err)
	}

	// Initialize routes
	router := routes.InitRoutes(database) // Assuming you might pass the database to init routes

	// Start the server
	startServer(router)
}

func startServer(router http.Handler) {
	srv := &http.Server{
		Addr:         configs.ServerAddress,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server is starting on port", configs.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	<-quit // Wait for interrupt signal

	log.Println("Shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
