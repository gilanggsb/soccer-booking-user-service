package models

import "time"

type Role struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Code      string `gorm:"varchar(50);not null"`
	Name      string `gorm:"varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
