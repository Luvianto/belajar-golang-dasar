package entity

import (
	commonModels "belajar-golang-dasar/common/models"

	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `gorm:"type:char(36);primaryKey;not null" validate:"required,uuid4"`
	IsAdmin  bool      `gorm:"not null"`
	Email    string    `gorm:"size:100;uniqueIndex;not null"`
	Password string    `gorm:"type:longtext;not null"`
	Phone    string    `gorm:"size:15;not null"`

	commonModels.TimestampsSoftDelete
}

type UserCreate struct {
	IsAdmin  bool   `json:"is_admin" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Phone    string `json:"phone" validate:"required,min=11,max=15"`
}

type UserReqByUUID struct {
	UUID string `uri:"uuid" binding:"required"`
}

type UserGet struct {
	UUID    uuid.UUID `json:"uuid"`
	IsAdmin bool      `json:"is_admin"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone" `
}
