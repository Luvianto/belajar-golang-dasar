package repository

import (
	memberEntity "belajar-golang-dasar/internal/module/member/entity"
	"belajar-golang-dasar/internal/module/member/interfaces"
	userEntity "belajar-golang-dasar/internal/module/user/entity"
	"belajar-golang-dasar/pkg/validator"

	"gorm.io/gorm"
)

var _ interfaces.MemberRepository = &memberRepository{}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{
		db: db,
	}
}

func (r *memberRepository) GetMember(id int) (*memberEntity.Member, bool, error) {
	var member memberEntity.Member
	query := r.db.Model(&member).Where("id = ?", id)
	exists, err := validator.Query(query)
	if err != nil {
		return nil, false, err
	}

	if !exists {
		return nil, false, nil
	}

	query.First(&member)
	return &member, true, nil
}

func (r *memberRepository) CreateMember(user userEntity.User, member memberEntity.Member) (*memberEntity.Member, bool, error) {
	userQuery := r.db.Model(&user).Create(&user)
	userExists, err := validator.Query(userQuery)
	if err != nil {
		return nil, false, err
	}

	if !userExists {
		return nil, false, nil
	}

	memberQuery := r.db.Model(&member).Create(&member)
	memberExists, err := validator.Query(memberQuery)
	if err != nil {
		return nil, false, err
	}

	if !memberExists {
		return nil, false, nil
	}

	return &member, true, nil
}
