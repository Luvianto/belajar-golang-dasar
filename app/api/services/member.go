package services

import (
	"belajar-golang-dasar/app/api/models"
	commonutils "belajar-golang-dasar/app/common/utils"
	"belajar-golang-dasar/app/config/repository"
)

type MemberService struct {
	Repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) *MemberService {
	return &MemberService{Repo: repo}
}

func (s *MemberService) GetAllMembers() ([]*models.Member, error) {
	return s.Repo.FetchAllMembers()
}

func (s *MemberService) GetMemberByID(id int) (*models.Member, error) {
	return s.Repo.FetchMemberByID(id)
}

func (s *MemberService) CreateMember(member *models.MemberCreate) (*models.Member, error) {
	userUUID := commonutils.GenerateUUID()
	return s.Repo.StoreMember(member, userUUID)
}

func (s *MemberService) UpdateMember(member *models.Member) (*models.Member, error) {
	return s.Repo.UpdateMember(member)
}

func (s *MemberService) DeleteMember(member *models.Member) (*models.Member, error) {
	return s.Repo.DestroyMember(member)
}
