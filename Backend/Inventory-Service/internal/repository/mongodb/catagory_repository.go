package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{
		collection: db.Collection("categories"),
	}
}

// Add your repository methods here, for example:
func (r *CategoryRepository) FindByID(ctx context.Context, id string) (*entity.Category, error) {
	// implementation
}
