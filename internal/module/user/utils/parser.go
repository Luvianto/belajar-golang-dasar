package utils

import (
	"belajar-golang-dasar/internal/module/user/entity"

	"github.com/google/uuid"
)

func UserCreateParser(user *entity.UserCreate, userUUID uuid.UUID) (*entity.User, error) {
	return &entity.User{
		UUID:     userUUID,
		IsAdmin:  false,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}, nil
}
