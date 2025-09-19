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

func (s *memberService) GetAllMember() ([]*memberEntity.MemberGet, error) {
	members, status, err := s.repo.GetAllMember()
	if err != nil {
		log.Error().Err(err).Msg("Gagal mendapatkan semua member")
		return nil, err
	}

	if !status {
		log.Error().Msg("Gagal mendapatkan semua member")
		return nil, fmt.Errorf("gagal mendapatkan semua member")
	}

	if len(members) == 0 {
		log.Error().Msg("Tidak ada member yang ditemukan")
		return nil, fmt.Errorf("tidak ada member yang ditemukan")
	}

	log.Info().Msg("Berhasil mendapatkan semua member")

	var memberGets []*memberEntity.MemberGet
	for _, member := range members {
		memberGets = append(memberGets, &memberEntity.MemberGet{
			ID: member.ID,
			User: userEntity.UserGet{
				UUID:    member.User.UUID,
				IsAdmin: member.User.IsAdmin,
				Email:   member.User.Email,
				Phone:   member.User.Phone,
			},
			Name:              member.Name,
			Major:             member.Major,
			ProfilePictureUrl: member.ProfilePictureUrl,
		})
	}

	return memberGets, nil
}

func (s *memberService) GetMember(req *memberEntity.MemberReqByID) (*memberEntity.MemberGet, error) {
	if req.ID == 0 {
		log.Error().Msg("ID tidak boleh kosong")
		return nil, fmt.Errorf("ID tidak boleh kosong")
	}

	member, found, err := s.repo.GetMember(req.ID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal mendapatkan member")
		return nil, err
	}

	if !found {
		log.Error().Msg("Member tidak ditemukan")
		return nil, fmt.Errorf("member tidak ditemukan")
	}

	log.Info().Msg("Berhasil mendapatkan member")

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
		log.Error().Err(err).Msg("Gagal memproses data member")
		return nil, err
	}

	createMember, err := memberUtils.MemberCreateParser(req, userUUID)
	if err != nil {
		log.Error().Err(err).Msg("Gagal memproses data member")
		return nil, err
	}

	member, status, err := s.repo.CreateMember(*createUser, *createMember)
	if err != nil {
		log.Error().Err(err).Msg("Gagal membuat member")
		return nil, err
	}

	if !status {
		log.Error().Msg("Member gagal dibuat")
		return nil, fmt.Errorf("member gagal dibuat")
	}

	log.Info().Msg("Berhasil membuat member")

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
		log.Error().Err(err).Msg("Gagal mengubah member")
		return nil, err
	}

	if !status {
		log.Error().Msg("Member gagal diubah")
		return nil, fmt.Errorf("member gagal diubah")
	}

	log.Info().Msg("Berhasil mengubah member")

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
		log.Error().Err(err).Msg("Gagal menghapus member")
		return nil, err
	}

	if !status {
		log.Error().Msg("Member gagal dihapus")
		return nil, fmt.Errorf("member gagal dihapus")
	}

	log.Info().Msg("Berhasil menghapus member")

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
