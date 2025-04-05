// internal/repository/refresh_token_repository.go
package repository

import (
	"fmt"
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/auth/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SaveRefreshToken ユーザーIDとリフレッシュトークンをデータベースに保存
func SaveRefreshToken(refreshtoken model.RefreshToken) error {
	if err := db.DB.Model(&model.RefreshToken{}).Create(refreshtoken).Error; err != nil {
		return err
	}
	return nil
}

// RefreshTokenExists トークンがデータベースに存在するか確認
func RefreshTokenExists(refreshToken string) (*model.RefreshToken, error) {
	var token model.RefreshToken
	if err := db.DB.Where("token = ?", refreshToken).First(&token).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("refresh token not found")
		}
		return nil, err
	}
	return &token, nil
}

// IsRefreshTokenValid トークンの有効期限を検証
func IsRefreshTokenValid(token *model.RefreshToken) bool {
	return token.Exp > time.Now().Unix()
}

// ValidateRefreshToken トークンの存在確認と有効期限の検証を統合
func ValidateRefreshToken(refreshToken string) (bool, uuid.UUID, error) {
	token, err := RefreshTokenExists(refreshToken)
	if err != nil {
		return false, uuid.Nil, err
	}

	if !IsRefreshTokenValid(token) {
		return false, uuid.Nil, fmt.Errorf("refresh token expired")
	}

	return true, token.UserID, nil
}

// DeleteRefreshTokenByUserID ユーザーIDに基づいてリフレッシュトークンを削除
// ユーザーがログアウトしたときに呼び出される
func DeleteRefreshTokenByUserID(userID uuid.UUID) error {
	var refreshToken model.RefreshToken
	if err := db.DB.Where("user_id = ?", userID).Delete(&refreshToken).Error; err != nil {
		return err
	}
	return nil
}
