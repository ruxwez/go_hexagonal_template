package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	authHandler "github.com/ruxwez/go_hexagonal_template/core/auth/handler"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {
	_authHandler := authHandler.NewHandler(services)

	app.Put("/auth", _authHandler.Register())
	app.Post("/auth", _authHandler.Login())
	app.Delete("/auth", _authHandler.Logout())
}
