package models

import (
	"time"

	
)

type Sale struct {
	ID         uint     `gorm:"primaryKey;autoIncrement"`
	CustomerID uint
	Customer   Customer
	SaleDate   time.Time
	Total      float64   `gorm:"type:numeric(12,2)"`
	CreatedBy  uint
	 User         User      `gorm:"foreignKey:CreatedBy"`
	UpdatedAt  time.Time

	Items []SaleItem `gorm:"foreignKey:SaleID"`
}
