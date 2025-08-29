package utils

import (
	"belajar-golang-dasar/app/api/models"

	"github.com/google/uuid"
)

func MemberCreateParser(member *models.MemberCreate, userUUID uuid.UUID) (*models.Member, error) {
	return &models.Member{
		UserID:            userUUID,
		Name:              member.Name,
		Major:             member.Major,
		ProfilePictureUrl: member.ProfilePictureUrl,
	}, nil
}
