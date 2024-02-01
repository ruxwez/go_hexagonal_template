package routes

import (
	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	authRoutes "github.com/ruxwez/go_hexagonal_template/core/auth/routes"
	userRoutes "github.com/ruxwez/go_hexagonal_template/core/user/routes"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {

	userRoutes.Setup(app, services)
	authRoutes.Setup(app, services)

}
