package db

import (
	"gorm.io/gorm"
)

// Database はデータベース操作のインターフェース
type Database interface {
	GetDB() *gorm.DB
}

// PostgresDatabase はPostgreSQLデータベース実装
type PostgresDatabase struct {
	db *gorm.DB
}

// NewPostgresDatabase は新しいPostgreSQLデータベースインスタンスを作成
func NewPostgresDatabase() (*PostgresDatabase, error) {
	db, err := InitPostgresDB()
	if err != nil {
		return nil, err
	}
	return &PostgresDatabase{db: db}, nil
}

// GetDB はGORMのDBインスタンスを返す
func (p *PostgresDatabase) GetDB() *gorm.DB {
	// 初期化済みのDBがある場合はそれを使用
	if p.db != nil {
		return p.db
	}
	// 後方互換性のため、グローバルDB変数にフォールバック
	return DB
}

// TestDatabase はテスト用データベース実装
type TestDatabase struct {
	db *gorm.DB
}

// NewTestDatabase は新しいテスト用データベースインスタンスを作成
func NewTestDatabase(db *gorm.DB) *TestDatabase {
	return &TestDatabase{db: db}
}

// GetDB はテスト用GORMのDBインスタンスを返す
func (t *TestDatabase) GetDB() *gorm.DB {
	return t.db
}
