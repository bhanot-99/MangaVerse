package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	r := gin.Default()

	// Product routes
	r.POST("/products", createProduct(db))
	r.GET("/products/:id", getProduct(db))
	r.PATCH("/products/:id", updateProduct(db))
	r.DELETE("/products/:id", deleteProduct(db))
	r.GET("/products", listProducts(db))

	log.Println("Inventory Service starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
}

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	r := gin.Default()

	// Product routes
	r.POST("/products", createProduct(db))
	r.GET("/products/:id", getProduct(db))
	r.PATCH("/products/:id", updateProduct(db))
	r.DELETE("/products/:id", deleteProduct(db))
	r.GET("/products", listProducts(db))

	log.Println("Inventory Service starting on :8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Create a new product
func createProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

// Get a product by ID
func getProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// Update a product by ID
func updateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// Delete a product by ID
func deleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Product{}, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}

// List all products
func listProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []Product
		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}
