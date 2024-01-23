package middlewares

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	"github.com/gofiber/fiber/v2"
)

func CheckToken(services *sharedApplication.Services) fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Next()
	}
}
