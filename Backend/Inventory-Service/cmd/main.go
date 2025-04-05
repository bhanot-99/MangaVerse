package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"inventory-service/internal/config"
	"inventory-service/internal/delivery/http"
	"inventory-service/internal/repository/mongodb"
	"inventory-service/internal/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	// MongoDB connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	db := client.Database(cfg.MongoDBName)

	// Initialize layers
	productRepo := mongodb.NewProductRepository(db)
	categoryRepo := mongodb.NewCategoryRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo, categoryRepo)

	// HTTP Server
	server := http.NewServer(productUsecase)
	go func() {
		if err := server.Run(cfg.HTTPPort); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	log.Printf("Inventory service running on port %s", cfg.HTTPPort)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
