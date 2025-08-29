package mysql

import (
	"belajar-golang-dasar/app/api/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchAllUsers() ([]*models.User, error) {
	var users []*models.User
	query := r.db.Find(&users)
	if query.Error != nil {
		return nil, query.Error
	}
	return users, nil
}

func (r *userRepository) FetchUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := r.db.Where("email = ?", email).First(&user)
	if query.Error != nil {
		return nil, query.Error
	}
	return &user, nil
}

func (r *userRepository) StoreUser(user *models.User) (*models.User, error) {
	query := r.db.Create(user)
	if query.Error != nil {
		return nil, query.Error
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	query := r.db.Save(user)
	if query.Error != nil {
		return nil, query.Error
	}
	return user, nil
}

func (r *userRepository) DestroyUser(user *models.User) (*models.User, error) {
	query := r.db.Delete(user)
	if query.Error != nil {
		return nil, query.Error
	}
	return user, nil
}
