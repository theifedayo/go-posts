package main

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/theifedayo/go-movie-api/config"
	"github.com/theifedayo/go-movie-api/api/models"
	"github.com/theifedayo/go-movie-api/config"
	"github.com/theifedayo/go-movie-api/api/routes"
	// import Redis and database drivers here
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	// Initialize the Redis cache
	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// Initialize the database connection
	db, err := models.ConnectDatabase(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// Initialize the API router
	router := routes.InitRoutes(db, cache)

	// Start the HTTP server
	log.Printf("Starting server on %s...", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
