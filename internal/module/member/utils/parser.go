package utils

import (
	"belajar-golang-dasar/internal/module/member/entity"

	"github.com/google/uuid"
)

func MemberCreateParser(member *entity.MemberCreate, userUUID uuid.UUID) (*entity.Member, error) {
	return &entity.Member{
		UserID:            userUUID,
		Name:              member.Name,
		Major:             member.Major,
		ProfilePictureUrl: member.ProfilePictureUrl,
	}, nil
}

func MemberUpdateParser(member *entity.MemberUpdate) (*entity.Member, error) {
	return &entity.Member{
		ID:                member.ID,
		UserID:            member.UserID,
		Name:              member.Name,
		Major:             member.Major,
		ProfilePictureUrl: member.ProfilePictureUrl,
	}, nil
}
