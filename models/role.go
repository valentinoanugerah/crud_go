package models

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(50);unique;not null"`

	Users []User `gorm:"foreignKey:RoleID"`
}
