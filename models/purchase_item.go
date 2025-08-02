package models

import "github.com/google/uuid"

type PurchaseItem struct {
	ID         uuid.UUID    `gorm:"primaryKey;autoIncrement"`
	PurchaseID uint
	Purchase   Purchase
	ProductID  uint
	Product    Product
	Quantity   int
	Price      float64 `gorm:"type:numeric(12,2)"`
}
