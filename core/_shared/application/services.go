package sharedApplication

import (
	authApplication "github.com/ruxwez/go_hexagonal_template/core/auth/application"
	userApplication "github.com/ruxwez/go_hexagonal_template/core/user/application"
)

type Services struct {
	UserService userApplication.UserService
	AuthService authApplication.AuthService
}
