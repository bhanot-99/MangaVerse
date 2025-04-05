package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Name        string                 `bson:"name"`
	Description string                 `bson:"description"`
	Price       float64                `bson:"price"`
	Stock       int                    `bson:"stock"`
	CategoryID  primitive.ObjectID     `bson:"category_id"`
	Metadata    map[string]interface{} `bson:"metadata,omitempty"` // For Jikan data
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
}
