package handler

import (
	"net/http"

	refresh "github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/refresh/service"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/service"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/auth"
	"github.com/gin-gonic/gin"
)

// RegisterHandler はユーザー登録を処理する
func RegisterHandler(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := service.RegisterUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func LoginHandler(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// ユーザー認証
	user, err := service.AuthenticateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 各種トークンを生成
	AccessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// internal/refresh/service/refresh.goに処理書いているけどおうだんするのどう？

	RefreshTokenStruct := refresh.SaveRefreshTokenStorage{}
	RefreshToken, err := auth.GenerateRefreshToken(user.ID, RefreshTokenStruct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// トークンを返す
	c.JSON(http.StatusOK, gin.H{
		"access_token":  AccessToken,
		"refresh_token": RefreshToken,
		"user":          gin.H{"id": user.ID, "email": user.Email},
	})
}

func ValidateAccessTokenHandler(c *gin.Context) {
	// Authorization ヘッダーからトークンを取得
	tokenString, err := auth.ExtractAccessTokenFromHeader(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
		c.Abort()
		return
	}
	// トークンの検証
	token, err := auth.ValidateAccessToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
