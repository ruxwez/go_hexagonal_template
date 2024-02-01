package userApplication

import (
	sharedDomain "github.com/ruxwez/go_hexagonal_template/core/_shared/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
)

type UserService interface {
	GetById(id int) *userDomain.User
}

type userService struct {
	user userDomain.UserRepository
}

func NewService(r *sharedDomain.Repositories) UserService {
	return &userService{
		user: r.UserRepository,
	}
}
