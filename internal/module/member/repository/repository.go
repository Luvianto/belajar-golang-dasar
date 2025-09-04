package repository

import (
	"belajar-golang-dasar/internal/module/member/entity"
	"belajar-golang-dasar/internal/module/member/interfaces"
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

func (r *memberRepository) GetMember(id int) (*entity.Member, bool, error) {
	var member entity.Member
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
