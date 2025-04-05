package auth

import (
	"crypto/rand"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// RefreshTokenStorage インターフェース
type SaveRefreshToken interface {
	SaveRefreshToken(userID uuid.UUID, refreshToken string) error
}
type ValidateRefreshToken interface {
	ValidateRefreshToken(refreshToken string) (bool, uuid.UUID, error)
}

// RefreshTokenを生成してハッシュ化した値を返す
func GenerateRefreshToken(userID uuid.UUID, saver SaveRefreshToken) (string, error) {
	// ランダムなリフレッシュトークンを生成
	refreshTokenBytes := make([]byte, 64) // 512-bitのリフレッシュトークン
	_, err := rand.Read(refreshTokenBytes)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}

	// ハッシュ化して保存（リフレッシュトークンは平文で保存しない）
	hashedRefreshToken, err := bcrypt.GenerateFromPassword(refreshTokenBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash refresh token")
	}
	err = saver.SaveRefreshToken(userID, string(hashedRefreshToken))
	if err != nil {
		return "", errors.Wrap(err, "failed to save refresh token")
	}

	return string(hashedRefreshToken), nil
}

// リフレッシュトークンを検証する
func RefreshAccessToken(refreshToken string, validator ValidateRefreshToken) (string, error) {
	isValid, userID, err := validator.ValidateRefreshToken(refreshToken)
	if err != nil || !isValid {
		return "", fmt.Errorf("invalid refresh token")
	}

	// 新しいアクセストークンの生成処理
	accessToken, err := GenerateAccessToken(userID)
	if err != nil {
		return "", fmt.Errorf("failed to generate access token: %v", err)
	}

	return accessToken, nil
}
