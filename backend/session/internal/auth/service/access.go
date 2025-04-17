package service

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/pkg/auth"
	"github.com/google/uuid"
)

// AccessTokenの生成
func GenerateAccessToken(userID uuid.UUID) (token string, exp int64, err error) {
	// ユーザーIDを使ってアクセストークンを生成
	accessToken, expiresAt, err := auth.GenerateAccessToken(userID)
	if err != nil {
		return "", int64(0), err
	}
	return accessToken, expiresAt, nil
}

// AccessTokenの検証
func ValidateAccessToken(token string) (bool, uuid.UUID, error) {
	// アクセストークンを検証
	userID, err := auth.GetUserIDFromToken(token)
	if err != nil {
		return false, uuid.Nil, err
	}
	return true, userID, nil
}
