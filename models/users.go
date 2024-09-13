package models

import (
	"time"
	_ "time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Password  string    `gorm:"column:password;not null"`
	Salt      string    `gorm:"column:salt;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
