<<<<<<< HEAD
// Create Order Handler
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	TotalPrice float64     `json:"total_price"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Order{}, &OrderItem{})

	r := gin.Default()

	// Order routes
	r.POST("/orders", createOrder(db))
	r.GET("/orders/:id", getOrder(db))
	r.PATCH("/orders/:id", updateOrder(db))
	r.GET("/orders", listOrders(db))

	log.Println("Order Service starting on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func createOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		// Calculate total price
		var totalPrice float64
		for _, item := range order.Items {
			totalPrice += item.Price * float64(item.Quantity)
		}
		order.TotalPrice = totalPrice

		// Create order in the database
		if err := db.Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(201, order)
	}
}

// Get Order Handler
func getOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order Order

		if err := db.Preload("Items").First(&order, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(200, order)
	}
}

// Update Order Handler
func updateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order Order

		// Find existing order
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		// Bind incoming JSON data to order struct
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		// Save updated order to database
		if err := db.Save(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update order"})
			return
		}

		c.JSON(200, order)
	}
}

// List Orders Handler
func listOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Order
		if err := db.Preload("Items").Find(&orders).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch orders"})
			return
		}

		c.JSON(200, orders)
	}
}
=======
// Create Order Handler
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	Status     string      `json:"status"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	TotalPrice float64     `json:"total_price"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func main() {
	// Initialize database
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Order{}, &OrderItem{})

	r := gin.Default()

	// Order routes
	r.POST("/orders", createOrder(db))
	r.GET("/orders/:id", getOrder(db))
	r.PATCH("/orders/:id", updateOrder(db))
	r.GET("/orders", listOrders(db))

	log.Println("Order Service starting on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func createOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		// Calculate total price
		var totalPrice float64
		for _, item := range order.Items {
			totalPrice += item.Price * float64(item.Quantity)
		}
		order.TotalPrice = totalPrice

		// Create order in the database
		if err := db.Create(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(201, order)
	}
}

// Get Order Handler
func getOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order Order

		if err := db.Preload("Items").First(&order, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(200, order)
	}
}

// Update Order Handler
func updateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var order Order

		// Find existing order
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "Order not found"})
			return
		}

		// Bind incoming JSON data to order struct
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(400, gin.H{"error": "Invalid data"})
			return
		}

		// Save updated order to database
		if err := db.Save(&order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update order"})
			return
		}

		c.JSON(200, order)
	}
}

// List Orders Handler
func listOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []Order
		if err := db.Preload("Items").Find(&orders).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch orders"})
			return
		}

		c.JSON(200, orders)
	}
}
>>>>>>> 8f5d13121b16ede815260d054d6c7ff5c19f075d
