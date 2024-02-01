package sharedDomain

import (
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
)

type Repositories struct {
	UserRepository userDomain.UserRepository
	AuthRepository authDomain.AuthRepository
}
