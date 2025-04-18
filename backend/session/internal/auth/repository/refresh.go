// internal/repository/refresh_token_repository.go
package repository

import (
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/internal/auth/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/pkg/db"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// SaveRefreshToken ある場合はリフレッシュトークンを更新し、なければ新規作成
func SaveRefreshToken(refreshtoken model.RefreshToken) error {
	if refreshtoken.UserID == uuid.Nil {
		return errors.New("invalid user ID")
	}
	if err := db.DB.Save(&refreshtoken).Error; err != nil {
		return err
	}
	return nil
}

// GetRefreshTokenModelFromUserID ユーザーIDからリフレッシュトークンを取得
func GetRefreshTokenModelFromUserID(userID uuid.UUID) (model.RefreshToken, error) {
	var refreshToken model.RefreshToken
	if err := db.DB.Where("user_id = ?", userID).First(&refreshToken).Error; err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}

// DeleteRefreshToken リフレッシュトークンでリフレッシュトークンを削除
func DeleteRefreshToken(userID uuid.UUID) error {
	if err := db.DB.Where("user_id = ?", userID).Delete(&model.RefreshToken{}).Error; err != nil {
		return err
	}
	return nil
}
