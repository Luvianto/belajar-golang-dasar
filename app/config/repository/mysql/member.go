package mysql

import (
	"belajar-golang-dasar/app/api/models"
	"belajar-golang-dasar/app/api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) FetchAllMembers() ([]*models.Member, error) {
	var members []*models.Member
	query := r.db.Find(&members)
	if query.Error != nil {
		return nil, query.Error
	}
	return members, nil
}

func (r *memberRepository) FetchMemberByID(id int) (*models.Member, error) {
	var member models.Member
	query := r.db.Where("id = ?", id).First(&member)
	if query.Error != nil {
		return nil, query.Error
	}
	return &member, nil
}

func (r *memberRepository) StoreMember(reqBody *models.MemberCreate, userUUID uuid.UUID) (*models.Member, error) {
	var returnMember *models.Member
	transaction := r.db.Transaction(func(tx *gorm.DB) error {
		user, err := utils.UserCreateParser(&reqBody.User, userUUID)
		if err != nil {
			return err
		}
		userQuery := tx.Create(user)
		if userQuery.Error != nil {
			return userQuery.Error
		}
		member, err := utils.MemberCreateParser(reqBody, userUUID)
		if err != nil {
			return err
		}
		memberQuery := tx.Create(member)
		if memberQuery.Error != nil {
			return memberQuery.Error
		}
		returnMember = member
		return nil
	})
	if transaction != nil {
		return nil, transaction
	}
	return returnMember, nil
}

func (r *memberRepository) UpdateMember(member *models.Member) (*models.Member, error) {
	query := r.db.Save(member)
	if query.Error != nil {
		return nil, query.Error
	}
	return member, nil
}

func (r *memberRepository) DestroyMember(member *models.Member) (*models.Member, error) {
	query := r.db.Delete(member)
	if query.Error != nil {
		return nil, query.Error
	}
	return member, nil
}
