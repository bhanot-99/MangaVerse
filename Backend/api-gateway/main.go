package main

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware for logging
	r.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Setup reverse proxies
	inventoryProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8081", // Proxy to inventory service
	})

	orderProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:8082", // Proxy to order service
	})

	// Routes for proxying requests
	r.Any("/inventory/*path", gin.WrapH(inventoryProxy))
	r.Any("/orders/*path", gin.WrapH(orderProxy))

	log.Println("API Gateway starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
