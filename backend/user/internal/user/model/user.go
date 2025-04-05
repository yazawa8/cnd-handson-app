package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`                              // UUID型の主キー
	Name         string    `gorm:"type:varchar(255)"`                                 // VARCHAR型
	Email        string    `gorm:"type:varchar(255);unique;not null"`                 // VARCHAR型 + UNIQUE制約
	PasswordHash string    `gorm:"type:text;not null"`                                // TEXT型
	RoleID       uuid.UUID `gorm:"type:uuid;not null"`                                // UUID型の外部キー
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp"` // TIMESTAMP型
	UpdatedAt    time.Time `gorm:"type:timestamp;not null"`                           // TIMESTAMP型
}
