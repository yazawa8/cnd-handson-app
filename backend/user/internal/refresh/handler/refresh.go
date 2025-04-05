package handler

import (
	"net/http"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/refresh/service"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/auth"
	"github.com/gin-gonic/gin"
)

func RefreshTokenHandler(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ValidateRefreshTokenStruct := service.ValidateRefreshTokenStorage{}
	newAccessToken, err := auth.RefreshAccessToken(req.RefreshToken, ValidateRefreshTokenStruct)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}
