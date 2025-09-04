package interfaces

import "belajar-golang-dasar/internal/module/user/entity"

type UserRepository interface {
	GetUser(id string) (*entity.User, bool, error)
}

type UserService interface {
	GetUser(req *entity.UserReqByUUID) (*entity.UserGet, error)
}
