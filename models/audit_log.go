package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID          uuid.UUID      `gorm:"primaryKey;autoIncrement"`
	UserID      uint
	User        User
	Action      string    `gorm:"type:varchar(100)"`
	TableName   string    `gorm:"type:varchar(100)"`
	RecordID    int
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
}
