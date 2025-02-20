package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommerceItem struct {
	ID         string  `gorm:"type:uuid;primaryKey"`
	ProductID  string  `gorm:"type:uuid" json:"product_id"`
	Quantity   int     `json:"quantity"`
	SiteID     string  `json:"site_id"`
	ListPrice  float64 `json:"list_price"`
	SalePrice  float64 `json:"sale_price"`
	Amount     float64 `json:"amount"`
	Discounted float64 `json:"discounted"`
	OrderID    string  `gorm:"type:uuid" json:"order_id"` // Chave estrangeira para Order
}

func (commerceItem *CommerceItem) BeforeCreate(tx *gorm.DB) (err error) {
	if commerceItem.ID == "" {
		commerceItem.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
