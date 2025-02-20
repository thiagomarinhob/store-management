package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Address     string `json:"address"`
	Number      string `json:"number"`
	References  string `json:"references"`
	Complement  string `json:"complement"`
	State       string `json:"state"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	FirstName   string `json:"first_name"`
	PhoneNumber string `json:"phone_number"`
	County      string `json:"county"`
	CustomerID  string `gorm:"type:uuid" json:"customer_id"` // Chave estrangeira para Customer
}

func (address *Address) BeforeCreate(tx *gorm.DB) (err error) {
	if address.ID == "" {
		address.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
