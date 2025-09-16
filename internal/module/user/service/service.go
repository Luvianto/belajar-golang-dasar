package service

import (
	"belajar-golang-dasar/internal/module/user/entity"
	"belajar-golang-dasar/internal/module/user/interfaces"
	"fmt"

	"github.com/rs/zerolog/log"
)

var _ interfaces.UserService = &userService{}

type userService struct {
	repo interfaces.UserRepository
}

func (s *userService) GetUser(req *entity.UserReqByUUID) (*entity.UserGet, error) {
	if req.UUID == "" {
		log.Error().Msg("UUID tidak boleh kosong")
		return nil, fmt.Errorf("UUID tidak boleh kosong")
	}

	user, found, err := s.repo.GetUser(req.UUID)
	if err != nil {
		log.Error().Msg("Gagal mendapatkan pengguna")
		return nil, fmt.Errorf("gagal mendapatkan pengguna")
	}

	if !found {
		log.Error().Msg("Pengguna tidak ditemukan")
		return nil, fmt.Errorf("pengguna tidak ditemukan")
	}

	return &entity.UserGet{
		UUID:    user.UUID,
		IsAdmin: user.IsAdmin,
		Email:   user.Email,
		Phone:   user.Phone,
	}, nil
}
