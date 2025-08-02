package models

import "github.com/google/uuid"

type SaleItem struct {
	ID        uuid.UUID    `gorm:"primaryKey;autoIncrement"`
	SaleID    uint
	Sale      Sale
	ProductID uint
	Product   Product
	Quantity  int
	Price     float64 `gorm:"type:numeric(12,2)"`
}
