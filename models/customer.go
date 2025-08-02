package model

import "github.com/google/uuid"

type Customer struct {
	ID      uuid.UUID   `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(100)"`
	Contact string `gorm:"type:varchar(100)"`

	Sales []Sale
}
