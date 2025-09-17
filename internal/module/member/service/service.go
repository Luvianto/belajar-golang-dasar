package service

import (
	commonutils "belajar-golang-dasar/common/utils"
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	"belajar-golang-dasar/internal/module/member/interfaces"
	memberUtils "belajar-golang-dasar/internal/module/member/utils"
	userEntity "belajar-golang-dasar/internal/module/user/entity"
	userUtils "belajar-golang-dasar/internal/module/user/utils"
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
		log.Error().Msg("ID tidak boleh kosong")
		return nil, fmt.Errorf("ID tidak boleh kosong")
	}

	member, found, err := s.repo.GetMember(req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal mendapatkan pengguna")
		return nil, err
	}

	if !found {
		log.Error().Msg("Pengguna tidak ditemukan")
		return nil, fmt.Errorf("pengguna tidak ditemukan")
	}

	log.Info().Msg("Berhasil mendapatkan pengguna")

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

func (s *memberService) CreateMember(req *memberEntity.MemberCreate) (*memberEntity.MemberGet, error) {
	if req == nil {
		log.Error().Msg("Member tidak boleh kosong")
		return nil, fmt.Errorf("member tidak boleh kosong")
	}

	userUUID := commonutils.GenerateUUID()

	createUser, err := userUtils.UserCreateParser(&req.User, userUUID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal memproses data pengguna")
		return nil, err
	}

	createMember, err := memberUtils.MemberCreateParser(req, userUUID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal memproses data member")
		return nil, err
	}

	member, status, err := s.repo.CreateMember(*createUser, *createMember)
	if err != nil {
		log.Error().Err(err).Msg("Gagal membuat pengguna")
		return nil, err
	}

	if !status {
		log.Error().Msg("Pengguna gagal dibuat")
		return nil, fmt.Errorf("pengguna gagal dibuat")
	}

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

func (s *memberService) UpdateMember(req *memberEntity.MemberUpdate) (*memberEntity.MemberGet, error) {
	if req == nil {
		log.Error().Msg("Member tidak boleh kosong")
		return nil, fmt.Errorf("member tidak boleh kosong")
	}

	updateMember, err := memberUtils.MemberUpdateParser(req)
	if err != nil {
		log.Error().Err(err).Msg("Gagal memproses data member")
		return nil, err
	}

	member, status, err := s.repo.UpdateMember(*updateMember)
	if err != nil {
		log.Error().Err(err).Msg("Gagal mengubah pengguna")
		return nil, err
	}

	if !status {
		log.Error().Msg("Pengguna gagal diubah")
		return nil, fmt.Errorf("pengguna gagal diubah")
	}

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

func (s *memberService) DeleteMember(req *memberEntity.MemberReqByID) (*memberEntity.MemberGet, error) {
	if req == nil {
		log.Error().Msg("Member tidak boleh kosong")
		return nil, fmt.Errorf("member tidak boleh kosong")
	}

	member, status, err := s.repo.DeleteMember(req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal menghapus pengguna")
		return nil, err
	}

	if !status {
		log.Error().Msg("Pengguna gagal dihapus")
		return nil, fmt.Errorf("pengguna gagal dihapus")
	}

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
