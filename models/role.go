package model

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(50);unique;not null"`

	Users []User `gorm:"foreignKey:RoleID"`
}
