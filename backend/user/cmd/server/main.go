package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/handler"
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
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

func runServer() {
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, err = db.InitDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	r := gin.Default()

	// 認証不要のエンドポイント（サインアップとログイン）
	r.POST("/user/register", handler.RegisterHandler)
	r.POST("/user/login", handler.LoginHandler)

	// 認証が必要なエンドポイント
	authorized := r.Group("/user")
	authorized.Use(auth.JWTMiddleware()) // ここでミドルウェア適用
	{
		// authorized.GET("/profile", handler.ProfileHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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

	err = db.MigrateDB(conn)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}

func loadEnv() error {
	return nil // 環境変数読み込みの実装（.env）
}
