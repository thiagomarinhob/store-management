package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Document    string    `json:"document"`
	DateOfBirth string    `json:"date_of_birth"`
	Telephone   string    `json:"telephone"`
	Active      bool      `json:"active"`
	StoreID     string    `gorm:"type:uuid" json:"store_id"`
	Addresses   []Address `gorm:"foreignKey:CustomerID"` // Especificando a chave estrangeira
	Orders      []Order   `gorm:"foreignKey:CustomerID"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	if customer.ID == "" {
		customer.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
