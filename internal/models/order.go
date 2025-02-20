package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID                string         `gorm:"type:uuid;primaryKey"`
	StoreID           string         `gorm:"type:uuid" json:"store_id"`
	SubmittedDate     time.Time      `json:"submitted_date"`
	State             string         `json:"state"`
	PaymentMethod     string         `json:"payment_method"`
	CommerceItems     []CommerceItem `gorm:"foreignKey:OrderID"` // Chave estrangeira para CommerceItem
	CustomerID        string         `gorm:"type:uuid" json:"customer_id"`
	ShippingAddressID string         `gorm:"type:uuid" json:"shipping_address_id"`
	ShippingAddress   Address        `gorm:"foreignKey:ID"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if order.ID == "" {
		order.ID = uuid.New().String() // Gerar o UUID se ele não existir
	}
	return nil
}
