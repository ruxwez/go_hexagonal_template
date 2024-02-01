package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	userHandler "github.com/ruxwez/go_hexagonal_template/core/user/handler"
	"github.com/ruxwez/go_hexagonal_template/middlewares"
)

func Setup(app *fiber.App, services *sharedApplication.Services) {
	_userHandler := userHandler.NewHandler(services)

	app.Get("/user", middlewares.CheckToken(services), _userHandler.GetDataByToken())
}
