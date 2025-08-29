package models

import (
	commonModels "belajar-golang-dasar/app/common/models"

	"github.com/google/uuid"
)

type Member struct {
	ID                int       `gorm:"primaryKey;autoIncrement;not null"`
	UserID            uuid.UUID `gorm:"not null"`
	Name              string    `gorm:"size:100;not null"`
	Major             string    `gorm:"not null"`
	ProfilePictureUrl string    `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	commonModels.TimestampsSoftDelete
}

type MemberCreate struct {
	User              UserCreate `json:"user" validate:"required"`
	Name              string     `json:"name" validate:"required,min=8,max=100"`
	Major             string     `json:"major" validate:"required"`
	ProfilePictureUrl string     `json:"profile_picture_url"`
}

type MemberResponse struct {
	User              User   `json:"user"`
	Name              string `json:"name"`
	Major             string `json:"major"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}
