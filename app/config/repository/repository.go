package repository

import (
	"belajar-golang-dasar/app/api/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	FetchAllUsers() ([]*models.User, error)
	FetchUserByEmail(email string) (*models.User, error)
	StoreUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DestroyUser(user *models.User) (*models.User, error)
}

type MemberRepository interface {
	FetchAllMembers() ([]*models.Member, error)
	FetchMemberByID(id int) (*models.Member, error)
	StoreMember(member *models.MemberCreate, userUUID uuid.UUID) (*models.Member, error)
	UpdateMember(member *models.Member) (*models.Member, error)
	DestroyMember(member *models.Member) (*models.Member, error)
}
