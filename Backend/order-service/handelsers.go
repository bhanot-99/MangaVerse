package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Start transaction
		tx := db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// Calculate total and prepare order items
		var total float64
		var orderItems []OrderItem

		for _, item := range req.Items {
			// Check product availability
			var product Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
				return
			}

			if product.Stock < item.Quantity {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product " + strconv.Itoa(int(item.ProductID))})
				return
			}

			// Update product stock
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
				return
			}

			// Add to order items
			itemTotal := float64(item.Quantity) * product.Price
			total += itemTotal

			orderItems = append(orderItems, OrderItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
		}

		// Create order
		order := Order{
			UserID: req.UserID,
			Status: OrderPending,
			Items:  orderItems,
			Total:  total,
		}

		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		// Commit transaction
		tx.Commit()

		c.JSON(http.StatusCreated, order)
	}
}

func getOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		var order Order
		if err := db.Preload("Items").First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func updateOrderStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
			return
		}

		var req UpdateOrderStatusRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var order Order
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		order.Status = req.Status
		if err := db.Save(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func getUserOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Query("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		var orders []Order
		query := db.Preload("Items").Where("user_id = ?", userID)

		// Filter by status if provided
		if status := c.Query("status"); status != "" {
			query = query.Where("status = ?", status)
		}

		// Pagination
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
		offset := (page - 1) * limit

		if err := query.Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}
