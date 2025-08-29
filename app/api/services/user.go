package services

import (
	"belajar-golang-dasar/app/api/models"
	"belajar-golang-dasar/app/config/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.Repo.FetchUserByEmail(email)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.Repo.FetchAllUsers()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.Repo.StoreUser(user)
}
