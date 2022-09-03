package gormrepository

import (
	"github.com/modhanami/dinkdonk/product-service/domain"
	"github.com/modhanami/dinkdonk/product-service/logger"
	"github.com/modhanami/dinkdonk/product-service/model/gormmodel"
	"github.com/modhanami/dinkdonk/product-service/service"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type productRepository struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewProductRepository(db *gorm.DB, logger logger.Logger) service.ProductRepository {
	return &productRepository{db, logger}
}

func (s *productRepository) GetProducts() ([]domain.Product, error) {
	var products []gormmodel.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}

	var result []domain.Product
	for _, p := range products {
		result = append(result, domain.Product{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			ImageURL:    p.ImageURL,
		})
	}

	return result, nil
}

func (s *productRepository) GetProductsWithCategories() ([]domain.Product, error) {
	var products []gormmodel.Product
	if err := s.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	var result []domain.Product
	for _, p := range products {
		var categories []domain.Category
		for _, c := range p.Categories {
			categories = append(categories, domain.Category(c.Name))
		}

		result = append(result, domain.Product{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: p.Description,
			ImageURL:    p.ImageURL,
			Categories:  categories,
		})
	}

	return result, nil
}
