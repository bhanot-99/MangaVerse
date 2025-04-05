package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		product := Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Stock:       req.Stock,
			Category:    req.Category,
			ImageURL:    req.ImageURL,
		}

		if err := db.Create(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

func getProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var product Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func updateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var product Product
		if err := db.First(&product, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		product.Name = req.Name
		product.Description = req.Description
		product.Price = req.Price
		product.Stock = req.Stock
		product.Category = req.Category
		product.ImageURL = req.ImageURL

		if err := db.Save(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func deleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		if err := db.Delete(&Product{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}

func listProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []Product
		query := db.Model(&Product{})

		// Filtering
		if category := c.Query("category"); category != "" {
			query = query.Where("category = ?", category)
		}

		// Search
		if search := c.Query("search"); search != "" {
			query = query.Where("name LIKE ?", "%"+search+"%")
		}

		// Pagination
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}
