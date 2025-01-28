package services

import (
	// Go Internal Packages
	"context"

	// Local Packages
	models "products/models"
)

type ProductRepository interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) error // UpdateProduct() error
	UpdateProduct(ctx context.Context, id int, product models.Product) error
	DeleteProductById(ctx context.Context, id int) error
}

type ProductService struct {
	repository ProductRepository
}

func NewService(repo ProductRepository) *ProductService {
	return &ProductService{repository: repo} // assign the repositoryInterface to the serviceClass
}

func (s *ProductService) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	return s.repository.GetProductById(ctx, id)
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.repository.GetAllProducts(ctx)
}

func (s *ProductService) CreateProduct(ctx context.Context, product models.Product) error {
	return s.repository.CreateProduct(ctx, product)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	return s.repository.UpdateProduct(ctx, id, product)
}

func (s *ProductService) DeleteProductById(ctx context.Context, id int) error {
	return s.repository.DeleteProductById(ctx, id)
}
