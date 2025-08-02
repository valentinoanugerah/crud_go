package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:numeric(12,2)"`
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time

	PurchaseItems []PurchaseItem
	SaleItems     []SaleItem
}
