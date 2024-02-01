package authApplication

import (
	sharedDomain "github.com/ruxwez/go_hexagonal_template/core/_shared/domain"
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
)

type AuthService interface {
	Register(user *userDomain.User, IP string) (string, error)
	Login(email string, password string, IP string) (string, error)
	Logout(token string) error
	CheckToken(token string) (int, error)
}

type authService struct {
	auth authDomain.AuthRepository
	user userDomain.UserRepository
}

func NewAuthService(r *sharedDomain.Repositories) AuthService {
	return &authService{
		auth: r.AuthRepository,
		user: r.UserRepository,
	}
}
