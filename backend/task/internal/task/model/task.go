package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`                              // UUID型の主キー
	title       string    `gorm:"type:varchar(255);not null"`                        // VARCHAR型
	description string    `gorm:"type:text;"`                                        // TEXT型
	status      string    `gorm:"type:varchar(255);not null"`                        // VARCHAR型
	start_time  time.Time `gorm:"type:timestamp;not null;default:current_timestamp"` // TIMESTAMP型
	end_time    time.Time `gorm:"type:timestamp;"`                                   // TIMESTAMP型
	column_id   uuid.UUID `gorm:"type:uuid;not null"`                                // UUID型の外部キー
	assignee_id uuid.UUID `gorm:"type:uuid;not null"`                                // UUID型の外部キー
}
