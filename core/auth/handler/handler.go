package authHandler

import (
	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	authApplication "github.com/ruxwez/go_hexagonal_template/core/auth/application"
)

type AuthHandler interface {
	Register() fiber.Handler
	Login() fiber.Handler
	Logout() fiber.Handler
}

type authHandler struct {
	authService authApplication.AuthService
}

func NewHandler(services *sharedApplication.Services) AuthHandler {
	return &authHandler{
		authService: services.AuthService,
	}
}
