package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100)"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	RoleID    uint
	Role      Role
	CreatedAt time.Time

	AuditLogs  []AuditLog  `gorm:"foreignKey:UserID"`
	Purchases  []Purchase  `gorm:"foreignKey:CreatedBy"`
	Sales        []Sale        `gorm:"foreignKey:CreatedBy"`
	Transactions []Transaction `gorm:"foreignKey:CreatedBy"`
}
