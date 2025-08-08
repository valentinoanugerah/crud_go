package models



type SaleItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	SaleID    uint
	Sale      Sale
	ProductID uint
	Product   Product
	Quantity  int
	Price     float64 `gorm:"type:numeric(12,2)"`
}
