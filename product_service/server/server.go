package server

import (
	"github.com/gin-gonic/gin"
	"github.com/modhanami/dinkdonk/product-service/domain"
	"github.com/modhanami/dinkdonk/product-service/logger"
)

type ProductService interface {
	GetProducts() ([]domain.Product, error)
	GetProductsWithCategories() ([]domain.Product, error)
}

type productHandler struct {
	service ProductService
	logger  logger.Logger
}

func NewProductHandler(service ProductService, logger logger.Logger) *productHandler {
	return &productHandler{service, logger}
}

//func (h *productHandler) GetProducts(c *gin.Context) {
//	products, err := h.service.GetProducts()
//	if err != nil {
//		c.JSON(500, gin.H{
//			"message": "failed to get products",
//		})
//		return
//	}
//
//	c.JSON(200, products)
//}

func (h *productHandler) GetProductsWithCategories(c *gin.Context) {
	products, err := h.service.GetProductsWithCategories()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get products",
		})
		h.logger.Error(err.Error())
		return
	}

	c.JSON(200, products)
}

func (h *productHandler) Mount(r *gin.Engine) {
	r.GET("/products", h.GetProductsWithCategories)
}
