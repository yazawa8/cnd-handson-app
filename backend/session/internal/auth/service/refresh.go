package service

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/internal/auth/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/session/internal/auth/repository"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// RefreshTokenの生成
func GenerateRefreshToken(userID uuid.UUID) (plainRefreshToken string, exp int64, err error) {
	// ランダムなリフレッシュトークンを生成
	refreshTokenBytes := make([]byte, 32) // 256-bitのリフレッシュトークン
	_, err = rand.Read(refreshTokenBytes)
	if err != nil {
		return "", int64(0), errors.New("failed to generate refresh token")
	}

	plainRefreshToken = base64.URLEncoding.EncodeToString(refreshTokenBytes)

	// ハッシュ化して保存（リフレッシュトークンは平文で保存しない）
	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(plainRefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", int64(0), errors.New("failed to hash refresh token")
	}

	expiresAt := time.Now().Add(30 * 24 * time.Hour).Unix() // 30日間有効
	refreshTokenModel := model.RefreshToken{
		UserID: userID,
		Token:  string(hashedRefreshToken),
		Exp:    expiresAt, // 30日間有効
	}

	err = repository.SaveRefreshToken(refreshTokenModel)
	if err != nil {
		return "", int64(0), errors.Wrap(err, "failed to save refresh token")
	}

	return plainRefreshToken, expiresAt, nil
}

// RefreshTokenの検証
func ValidateRefreshToken(userID uuid.UUID, plainRefreshToken string) error {
	// データベースからハッシュ化されたトークンを取得
	storedToken, err := repository.GetRefreshTokenModelFromUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve refresh token")
	}

	// クライアントから送信されたトークンを検証
	err = bcrypt.CompareHashAndPassword([]byte(storedToken.Token), []byte(plainRefreshToken))
	if err != nil {
		return errors.New("invalid refresh token")
	}

	return nil
}

// RefreshTokenの無効化
func RevokeRefreshToken(userID uuid.UUID, plainRefreshToken string) error {
	// リフレッシュトークンをハッシュ化して保存されたものと比較
	_, err := bcrypt.GenerateFromPassword([]byte(plainRefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "failed to hash refresh token")
	}
	err = repository.DeleteRefreshToken(userID)
	if err != nil {
		return errors.Wrap(err, "failed to delete refresh token")
	}
	return nil
}
