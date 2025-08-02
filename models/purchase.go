package models

import (
	"time"

	"github.com/google/uuid"
)

type Purchase struct {
	ID           uuid.UUID      `gorm:"primaryKey;autoIncrement"`
	SupplierID   uint
	Supplier     Supplier
	PurchaseDate time.Time
	Total        float64   `gorm:"type:numeric(12,2)"`
	CreatedBy    uint
	User         User      `gorm:"foreignKey:CreatedBy"` 

	Items []PurchaseItem `gorm:"foreignKey:PurchaseID"`
}
