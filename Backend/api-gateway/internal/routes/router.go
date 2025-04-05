package routes

import (
	"github.com/gin-gonic/gin" // Keep external dependencies as-is

	"api-gateway/internal/config"
	"api-gateway/internal/handlers"
	"api-gateway/internal/middleware"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()

	// Middleware
	router.Use(middleware.Logger(cfg))
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Health check
	router.GET("/health", handlers.HealthCheck)

	// API routes
	api := router.Group("/api")
	{
		// Inventory service routes
		inventory := api.Group("/products")
		inventory.Use(middleware.Authenticate(cfg.JWTSecret))
		{
			inventory.GET("", handlers.ProxyToInventoryService(cfg))
			inventory.GET("/:id", handlers.ProxyToInventoryService(cfg))
			inventory.POST("", handlers.ProxyToInventoryService(cfg))
			inventory.PATCH("/:id", handlers.ProxyToInventoryService(cfg))
			inventory.DELETE("/:id", handlers.ProxyToInventoryService(cfg))
		}

		// Order service routes
		orders := api.Group("/orders")
		orders.Use(middleware.Authenticate(cfg.JWTSecret))
		{
			orders.GET("", handlers.ProxyToOrderService(cfg))
			orders.GET("/:id", handlers.ProxyToOrderService(cfg))
			orders.POST("", handlers.ProxyToOrderService(cfg))
			orders.PATCH("/:id", handlers.ProxyToOrderService(cfg))
		}
	}

	return router
}
