package service

import (
	"github.com/modhanami/dinkdonk/product-service/domain"
	"github.com/modhanami/dinkdonk/product-service/server"
)

type ProductRepository interface {
	GetProducts() ([]domain.Product, error)
	GetProductsWithCategories() ([]domain.Product, error)
}

type productService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) server.ProductService {
	return &productService{repo}
}

func (s *productService) GetProducts() ([]domain.Product, error) {
	return s.repo.GetProducts()
}

func (s *productService) GetProductsWithCategories() ([]domain.Product, error) {
	return s.repo.GetProductsWithCategories()
}
