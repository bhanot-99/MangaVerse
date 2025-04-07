package usecase

import (
	"context"
	"errors"

	"inventory-service/internal/repository/mongodb"
)

type ProductUsecase struct {
	productRepo  mongodb.ProductRepository
	categoryRepo mongodb.CategoryRepository
}

func NewProductUsecase(productRepo mongodb.ProductRepository, categoryRepo mongodb.CategoryRepository) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

// CreateProduct remains the same as existing implementation
func (uc *ProductUsecase) CreateProduct(ctx context.Context, product Product) (*Product, error) {
	// Validate category exists
	if _, err := uc.categoryRepo.FindByID(ctx, product.CategoryID); err != nil {
		return nil, errors.New("category does not exist")
	}

	createdProduct, err := uc.productRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return createdProduct, nil
}

// GetProductByID remains the same as existing implementation
func (uc *ProductUsecase) GetProductByID(ctx context.Context, id string) (*Product, error) {
	return uc.productRepo.FindByID(ctx, id)
}

// UpdateProduct updates an existing product
func (uc *ProductUsecase) UpdateProduct(ctx context.Context, id string, product Product) error {
	// Validate product exists
	if _, err := uc.GetProductByID(ctx, id); err != nil {
		return errors.New("product not found")
	}

	// Validate category exists if being updated
	if product.CategoryID != "" {
		if _, err := uc.categoryRepo.FindByID(ctx, product.CategoryID); err != nil {
			return errors.New("category does not exist")
		}
	}

	return uc.productRepo.Update(ctx, id, product)
}

// DeleteProduct removes a product
func (uc *ProductUsecase) DeleteProduct(ctx context.Context, id string) error {
	// Validate product exists
	if _, err := uc.GetProductByID(ctx, id); err != nil {
		return errors.New("product not found")
	}

	return uc.productRepo.Delete(ctx, id)
}

// ListProducts returns paginated list of products
func (uc *ProductUsecase) ListProducts(ctx context.Context, page, limit int) ([]Product, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	return uc.productRepo.FindAll(ctx, page, limit)
}

// Product represents the domain model
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CategoryID  string  `json:"category_id"`
}
