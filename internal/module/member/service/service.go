package service

import (
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	"belajar-golang-dasar/internal/module/member/interfaces"
	userEntity "belajar-golang-dasar/internal/module/user/entity"
	"fmt"

	"github.com/rs/zerolog/log"
)

var _ interfaces.MemberService = &memberService{}

type memberService struct {
	repo interfaces.MemberRepository
}

func NewMemberService(repo interfaces.MemberRepository) *memberService {
	return &memberService{
		repo: repo,
	}
}

func (s *memberService) GetMember(req *memberEntity.MemberReqByID) (*memberEntity.MemberGet, error) {
	if req.ID == 0 {
		log.Error().Msg("id tidak boleh kosong")
		return nil, fmt.Errorf("id tidak boleh kosong")
	}

	member, found, err := s.repo.GetMember(req.ID)
	if err != nil {
		log.Error().Err(err).Msg("gagal mendapatkan pengguna")
		return nil, err
	}
	if !found {
		log.Error().Msg("pengguna tidak ditemukan")
		return nil, fmt.Errorf("pengguna tidak ditemukan")
	}

	log.Info().Msg("berhasil mendapatkan pengguna")
	return &memberEntity.MemberGet{
		ID: member.ID,
		User: userEntity.UserGet{
			UUID:    member.UserID,
			IsAdmin: member.User.IsAdmin,
			Email:   member.User.Email,
			Phone:   member.User.Phone,
		},
		Name:              member.Name,
		Major:             member.Major,
		ProfilePictureUrl: member.ProfilePictureUrl,
	}, nil
}
