package commonModels

import (
	"time"

	"gorm.io/gorm"
)

type TimestampsSoftDelete struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
