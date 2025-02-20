package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Media struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string `json:"name"`
	Path string `json:"path"`
}

func (media *Media) BeforeCreate(tx *gorm.DB) (err error) {
	if media.ID == "" {
		media.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
