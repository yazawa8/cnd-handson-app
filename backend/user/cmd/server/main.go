package main

import (
	"fmt"
	"log"
	"os"

	authHandler "github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/auth/handler"
	refreshModel "github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/auth/model"
	userModel "github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/model"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/auth"
	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: myapp [server|migrate]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "server":
		runServer()
	case "migrate":
		runMigrate()
	case "reset":
		resetDB()
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func runServer() {
	// init処理
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, err = db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// start server
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// not authenticated endpoints
	r.POST("/auth/register", authHandler.RegisterHandler)
	r.POST("/auth/login", authHandler.LoginHandler)
	r.GET("/auth/validate", authHandler.ValidateAccessTokenHandler)
	r.POST("/auth/refresh", authHandler.RefreshTokenHandler)

	// authenticated endpoints
	r.POST("/auth/logout", auth.JWTMiddleware(), authHandler.LogoutHandler)
	authorized := r.Group("/user")
	authorized.Use(auth.JWTMiddleware()) // ここでミドルウェア適用
	{
		// authorized.GET("/profile", handler.ProfileHandler)

	}

	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func runMigrate() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.MigrateDB(conn, userModel.User{}, refreshModel.RefreshToken{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}
func resetDB() {
	conn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.ResetDB(conn, userModel.User{}, refreshModel.RefreshToken{})
	if err != nil {
		log.Fatalf("Database reset failed: %v", err)
	}

	log.Println("Database reset completed successfully.")
}

func loadEnv() error {
	return nil // 環境変数読み込みの実装（.env）
}
