package utils

import (
	"belajar-golang-dasar/app/api/models"
	commonutils "belajar-golang-dasar/app/common/utils"

	"github.com/google/uuid"
)

func UserCreateParser(user *models.UserCreate, userUUID uuid.UUID) (*models.User, error) {
	encryptedPassword, err := commonutils.Encrypt(&user.Password)
	if err != nil {
		return nil, err
	}

	return &models.User{
		UUID:     userUUID,
		IsAdmin:  user.IsAdmin,
		Email:    user.Email,
		Password: encryptedPassword,
		Phone:    user.Phone,
	}, nil
}
