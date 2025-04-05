package auth

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// ユーザーIDからアクセストークンを生成する
func GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign JWT")
	}

	return tokenString, nil
}

// リクエストヘッダーからトークンを抽出する
func ExtractAccessTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("Authorization header is required")
	}

	// "Bearer <token>" の形式か確認
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid Authorization format")
	}

	return parts[1], nil
}

// アクセストークンを検証する
func ValidateAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	return token, err
}

// アクセストークンからユーザーIDを取得する
func GetUserIDFromToken(tokenString string) (uuid.UUID, error) {
	token, err := ValidateAccessToken(tokenString)
	if err != nil {
		return uuid.Nil, errors.Wrap(err, "failed to validate token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.New("invalid token claims")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return uuid.Nil, errors.New("user ID not found in token claims")
	}

	return uuid.Parse(userID)
}
