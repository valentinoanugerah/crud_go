package models

type PurchaseItem struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	PurchaseID uint
	Purchase   Purchase
	ProductID  uint
	Product    Product
	Quantity   int
	Price      float64 `gorm:"type:numeric(12,2)"`
}
