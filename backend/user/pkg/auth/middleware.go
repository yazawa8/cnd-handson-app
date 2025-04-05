// pkg/auth/middleware.go
package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware は JWT 認証を行うミドルウェア
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization ヘッダーからトークンを取得
		tokenString, err := ExtractAccessTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			c.Abort()
			return
		}

		// トークンの検証
		token, err := ValidateAccessToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// トークンが有効ならリクエストを続行
		c.Next()
	}
}
