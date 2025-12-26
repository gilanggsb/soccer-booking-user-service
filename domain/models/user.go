package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Username    string    `gorm:"type:varchar(255);not null"`
	Password    string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	RoleID      uint      `gorm:"type:uint;not null"`
	PhoneNumber string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Role        Role `gorm:"foreignKey:role_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
