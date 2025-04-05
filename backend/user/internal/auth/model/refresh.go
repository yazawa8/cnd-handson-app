package model

import "github.com/google/uuid"

// RefreshToken リフレッシュトークンの構造体
type RefreshToken struct {
	UserID uuid.UUID `json:"user_id"`
	Token  string    `json:"token"`
	Exp    int64     `json:"expired_at"`
}
