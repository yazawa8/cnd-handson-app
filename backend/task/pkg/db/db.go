package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init 関数でデータベース接続を確立
func init() {
	// 初回接続の試行
	connectDB()
}

// connectDB はデータベースに接続する処理を担当
func connectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DB")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)

	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	DB = db
}

// InitDB は接続が必要な場合に再接続を試みる
func InitDB() (*gorm.DB, error) {
	log.Println("Checking database connection...")
	if DB == nil || !isDBConnected() {
		log.Println("Database connection lost. Retrying...")
		connectDB() // 再接続
	}

	// 接続状態を再確認
	if DB == nil || !isDBConnected() {
		return nil, fmt.Errorf("failed to reconnect to the database")
	}

	return DB, nil
}

// isDBConnected はデータベースが接続されているか確認する関数
func isDBConnected() bool {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return false
	}
	if err := sqlDB.Ping(); err != nil {
		return false
	}
	return true
}
