package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID          string  `gorm:"type:uuid;primaryKey"`
	Name        string  `json:"name"`
	AddressID   *string `gorm:"type:uuid" json:"address_id"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Admins      []Admin `gorm:"foreignKey:StoreID"`
}

func (store *Store) BeforeCreate(tx *gorm.DB) (err error) {
	if store.ID == "" {
		store.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
