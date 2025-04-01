package db

import (
	"log"

	"github.com/cloudnativedaysjp/cnd-handson-app/backend/user/internal/user/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	log.Println("Starting migration...")

	// テーブルのマイグレーション
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	log.Println("Migration completed successfully.")
	return nil
}
