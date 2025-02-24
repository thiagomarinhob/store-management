package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID       string  `gorm:"type:uuid;primaryKey"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	StoreID  *string `gorm:"type:uuid" json:"store_id"`
}

func (admin *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if admin.ID == "" {
		admin.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
