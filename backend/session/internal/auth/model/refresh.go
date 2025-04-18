package model

import "github.com/google/uuid"

// RefreshToken リフレッシュトークンの構造体
type RefreshToken struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Token  string    `gorm:"not null"`
	Exp    int64     `gorm:"not null"`
}
