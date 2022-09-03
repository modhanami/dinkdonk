package gormmodel

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"not null;"`
	Price       float64 `gorm:"not null;"`
	Description string
	ImageURL    string
	Categories  []Category `gorm:"many2many:product_categories;"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"not null;"`
}

//type Category struct {
//	ID         uint `gorm:"primarykey"`
//	ProductID  uint `gorm:"not null"`
//	CategoryID uint `gorm:"not null"`
//	Category   Category
//}
