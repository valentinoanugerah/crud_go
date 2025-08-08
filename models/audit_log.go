package models

import (
	"time"


)

type AuditLog struct {
	ID          uint     `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	User        User
	Action      string    `gorm:"type:varchar(100)"`
	TableName   string    `gorm:"type:varchar(100)"`
	RecordID    int
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
}
