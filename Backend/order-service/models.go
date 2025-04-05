package main

import "gorm.io/gorm"

type OrderStatus string

const (
	OrderPending   OrderStatus = "pending"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	gorm.Model
	UserID uint        `json:"user_id" gorm:"not null"`
	Status OrderStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	Items  []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	Total  float64     `json:"total" gorm:"not null"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Quantity  int     `json:"quantity" gorm:"not null;check:quantity > 0"`
	Price     float64 `json:"price" gorm:"not null;check:price >= 0"`
}

type CreateOrderRequest struct {
	UserID uint               `json:"user_id" binding:"required"`
	Items  []OrderItemRequest `json:"items" binding:"required,min=1"`
}

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status" binding:"required,oneof=pending completed cancelled"`
}

// Product model needed for order processing
type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null;check:price >= 0"`
	Stock       int     `json:"stock" gorm:"not null;check:stock >= 0"`
	Category    string  `json:"category" gorm:"not null"`
	ImageURL    string  `json:"image_url"`
}
