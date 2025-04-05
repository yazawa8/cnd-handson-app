package service

import (
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/refresh/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/refresh/repository"
	"github.com/google/uuid"
)

// SaveRefreshTokenStorage構造体
type SaveRefreshTokenStorage struct{}
type ValidateRefreshTokenStorage struct{}

// SaveRefreshToken ユーザーIDとリフレッシュトークンをデータベースに保存
func (s SaveRefreshTokenStorage) SaveRefreshToken(userID uuid.UUID, refreshToken string) error {
	refreshTokenModel := model.RefreshToken{
		UserID: userID,
		Token:  refreshToken,
		Exp:    time.Now().Add(time.Hour * 24 * 30).Unix(), // 30日間有効
	}

	return repository.SaveRefreshToken(refreshTokenModel)
}

// ValidateRefreshToken リフレッシュトークンがデータベースに存在するか、かつ有効かを確認
func (s ValidateRefreshTokenStorage) ValidateRefreshToken(refreshToken string) (bool, uuid.UUID, error) {
	return repository.ValidateRefreshToken(refreshToken)
}
