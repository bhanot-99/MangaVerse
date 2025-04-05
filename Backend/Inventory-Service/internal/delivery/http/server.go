package http

import (
	"context"
	"inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router         *gin.Engine
	productUsecase usecase.ProductUsecase
}

func NewServer(productUsecase usecase.ProductUsecase) *Server {
	s := &Server{
		router:         gin.New(),
		productUsecase: productUsecase,
	}

	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	s.setupRoutes()

	return s
}

func (s *Server) Run(port string) error {
	return s.router.Run(":" + port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	// Add any cleanup logic here
	return nil
}

func (s *Server) setupRoutes() {
	v1 := s.router.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", s.createProduct)
			products.GET("/:id", s.getProduct)
			products.PATCH("/:id", s.updateProduct)
			products.DELETE("/:id", s.deleteProduct)
			products.GET("", s.listProducts)
		}

		// Add categories routes if needed
	}
}

// Implement handler methods (createProduct, getProduct, etc.) below
