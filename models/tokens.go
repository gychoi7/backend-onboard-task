package models

import (
	"time"
	_ "time"
)

type Token struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"column:user_id;not null;constraint:OnDelete:CASCADE;"`
	Refresh   string    `gorm:"column:refresh;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
