package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null;check:price >= 0"`
	Stock       int     `json:"stock" gorm:"not null;check:stock >= 0"`
	Category    string  `json:"category" gorm:"not null"`
	ImageURL    string  `json:"image_url"`
}

// ProductRequest is used for create/update operations
type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url"`
}
