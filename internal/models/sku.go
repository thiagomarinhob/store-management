package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sku struct {
	ID         string            `gorm:"type:uuid;primaryKey"`
	SkuName    string            `json:"sku_name"`
	Attributes map[string]string `gorm:"type:jsonb" json:"attributes"`
	ProductID  string            `gorm:"type:uuid" json:"product_id"`
}

func (sku *Sku) BeforeCreate(tx *gorm.DB) (err error) {
	if sku.ID == "" {
		sku.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
