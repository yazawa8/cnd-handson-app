package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// .envファイルを読み込む
	loadEnvFile()
	
	// 初回接続の試行
	connectDB()
}

// loadEnvFile は.envファイルを読み込む
func loadEnvFile() {
	// カレントディレクトリから親ディレクトリへ順番に.envファイルを検索
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get current directory: %v", err)
		return
	}
	
	// 最大3階層まで遡ってenvファイルを探す
	for i := 0; i < 3; i++ {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			if err := godotenv.Load(envPath); err == nil {
				log.Printf("Loaded .env file from: %s", envPath)
				return
			}
		}
		// 親ディレクトリへ
		dir = filepath.Dir(dir)
	}
	
	log.Println("Warning: .env file not found")
}

// connectDB はデータベースに接続する処理を担当
func connectDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DB")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	
	log.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// エラーログを出力するだけでパニックは発生させない
		log.Printf("Failed to connect to the database: %v", err)
		return
	}
	log.Println("Successfully connected to database")
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
		return nil, fmt.Errorf("Failed to reconnect to the database")
	}

	return DB, nil
}

// isDBConnected はデータベースが接続されているか確認する関数
func isDBConnected() bool {
	// DBがnilの場合は接続されていない
	if DB == nil {
		return false
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get database connection: %v", err)
		return false
	}
	if err := sqlDB.Ping(); err != nil {
		log.Printf("Failed to ping database: %v", err)
		return false
	}
	return true
}