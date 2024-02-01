package userHandler

import (
	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	userApplication "github.com/ruxwez/go_hexagonal_template/core/user/application"
)

type UserHandler interface {
	GetDataByToken() fiber.Handler
}

type userHandler struct {
	userService userApplication.UserService
}

func NewHandler(services *sharedApplication.Services) UserHandler {
	return &userHandler{
		userService: services.UserService,
	}
}
