package model

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`                              // UUID型の主キー
	Name        string    `gorm:"type:varchar(255);not null"`                        // VARCHAR型
	Description string    `gorm:"type:text"`                                         // TEXT型
	OwnerID     uuid.UUID `gorm:"type:uuid;not null"`                                // UUID型の外部キー（オーナー）
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:current_timestamp"` // TIMESTAMP型
	UpdatedAt   time.Time `gorm:"type:timestamp;not null"`                           // TIMESTAMP型
}