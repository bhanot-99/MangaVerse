package usecase

import (
	"context"
	"errors"

	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
)

type ProductUsecase struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductUsecase(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *ProductUsecase) CreateProduct(ctx context.Context, input entity.Product) (*entity.Product, error) {
	// Validate input
	if input.Name == "" {
		return nil, errors.New("product name is required")
	}
	if input.Price <= 0 {
		return nil, errors.New("price must be positive")
	}

	return uc.productRepo.Create(ctx, &input)
}

func (uc *ProductUsecase) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	return uc.productRepo.GetByID(ctx, id)
}

// Implement other use cases (Update, Delete, List) similarly
