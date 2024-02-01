package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	"github.com/ruxwez/go_hexagonal_template/core/_shared/domain/response"
)

func CheckToken(services *sharedApplication.Services) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userId int
		var err error

		// Obtenemos el header de Autorización
		authToken := c.Get("Authorization")

		// Validamos el token
		if userId, err = services.AuthService.CheckToken(authToken); err != nil {
			return c.Status(http.StatusUnauthorized).JSON(response.Error{
				Error: err.Error(),
			})
		}

		c.Locals("userID", userId)

		return c.Next()
	}
}
