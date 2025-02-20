package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	Media       []Media   `gorm:"many2many:category_media"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
	StoreID     string    `gorm:"type:uuid" json:"store_id"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if category.ID == "" {
		category.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
