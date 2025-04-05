package db

import (
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB, models ...interface{}) error {
	log.Println("Starting migration...")

	// テーブルのマイグレーション
	err := db.AutoMigrate(models...)
	if err != nil {
		return err
	}

	log.Println("Migration completed successfully.")
	return nil
}

func ResetDB(db *gorm.DB, models ...interface{}) error {
	log.Println("Resetting database...")

	// 既存のテーブルを削除
	for _, model := range models {
		err := db.Migrator().DropTable(model)
		log.Println("Dropping table:", model)
		if err != nil {
			return err
		}
	}
	log.Println("All tables dropped successfully.")

	// テーブルを再作成
	err := MigrateDB(db, models...)
	if err != nil {
		return err
	}

	log.Println("Database reset completed successfully.")
	return nil
}
