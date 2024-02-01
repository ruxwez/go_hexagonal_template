package authHandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ruxwez/go_hexagonal_template/core/_shared/domain/response"
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	tokenUtils "github.com/ruxwez/go_hexagonal_template/utils/token"
)

func (a *authHandler) Logout() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		// Comprobamos el token
		token, ok := tokenUtils.ValidateToken(token)
		if !ok {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: authDomain.ErrEmptyToken.Error(),
			})
		}

		err := a.authService.Logout(token)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(response.Error{
				Error: err.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(response.Default{
			Message: "Sesión cerrada correctamente",
		})
	}
}
