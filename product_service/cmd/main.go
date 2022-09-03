package main

import (
	"github.com/gin-gonic/gin"
	"github.com/modhanami/dinkdonk/product-service/logger/zaplogger"
	"github.com/modhanami/dinkdonk/product-service/model/gormmodel"
	"github.com/modhanami/dinkdonk/product-service/repository/gormrepository"
	"github.com/modhanami/dinkdonk/product-service/server"
	"github.com/modhanami/dinkdonk/product-service/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	db, err := initGormDB()
	if err != nil {
		return err
	}
	err = createFakeProducts(db)
	if err != nil {
		return err
	}

	l := zaplogger.NewZapLogger()
	repo := gormrepository.NewProductRepository(db, nil)
	svc := service.NewProductService(repo)
	handler := server.NewProductHandler(svc, l)

	handler.Mount(r)

	l.Info("starting product service")

	return r.Run("localhost:30027")
}

func initGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&gormmodel.Product{}, &gormmodel.Category{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createFakeProducts(db *gorm.DB) error {
	categories := []gormmodel.Category{
		{
			Name: "Category 1",
		},
		{
			Name: "Category 2",
		},
		{
			Name: "Category 3",
		},
	}

	if err := db.Create(&categories).Error; err != nil {
		return err
	}

	products := []gormmodel.Product{
		{
			Name:        "User 1",
			Price:       100,
			Description: "This is product 1",
			ImageURL:    "https://example.com/product1.png",
			Categories:  []gormmodel.Category{categories[0], categories[1]},
		},
		{
			Name:        "User 2",
			Price:       200,
			Description: "This is product 2",
			ImageURL:    "https://example.com/product2.png",
			Categories:  categories,
		},
		{
			Name:        "User 3",
			Price:       300,
			Description: "This is product 3",
			ImageURL:    "https://example.com/product3.png",
		},
	}

	return db.Create(&products).Error
}
