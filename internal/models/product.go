package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `gorm:"type:uuid;primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	ListPrice   float64 `json:"list_price"`
	SalePrice   float64 `json:"sale_price"`
	Active      bool    `json:"active"`
	SKUs        []Sku   `gorm:"foreignKey:ProductID"`
	CategoryID  string  `gorm:"type:uuid" json:"category_id"`
	StoreID     string  `gorm:"type:uuid" json:"store_id"`
	Media       []Media `gorm:"many2many:product_media"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	if product.ID == "" {
		product.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
