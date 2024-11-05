package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
}
