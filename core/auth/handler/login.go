package authHandler

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	sharedDomain "github.com/ruxwez/go_hexagonal_template/core/_shared/domain"
	"github.com/ruxwez/go_hexagonal_template/core/_shared/domain/response"
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	emailUtils "github.com/ruxwez/go_hexagonal_template/utils/email"
)

func (u *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Creamos el modelo con la estructura del body
		bodyModel := authDomain.LoginBodyRequest{}

		// Parseamos el body de la peticion
		if err := c.BodyParser(&bodyModel); err != nil {
			// En caso de error, devolvemos un error 400 y el codigo de error
			return c.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: sharedDomain.ErrBodyNotValid.Error(),
			})
		}

		// Limpiamos los datos del modelo
		bodyModel.Email = strings.TrimSpace(bodyModel.Email)
		bodyModel.Password = strings.TrimSpace(bodyModel.Password)

		// Comprobamos que el email no este vacio
		if bodyModel.Email == "" {
			return c.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmptyEmail.Error(),
			})
		}

		// Comprobamos que el email sea valido
		if !emailUtils.Check(bodyModel.Email) {
			return c.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmailNotValid.Error(),
			})
		}

		// Comprobamos que la contraseña no este vacia, tenga al menos 8 caracteres
		if bodyModel.Password == "" || len(bodyModel.Password) < 8 {
			return c.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrPasswordNotValid.Error(),
			})
		}

		authToken, err := u.authService.Login(bodyModel.Email, bodyModel.Password, c.IP())

		// Comprobamos si ha habido algun error
		if err != nil {
			// En caso de error, devolvemos un error 400 y el codigo de error que nos devuelve el servicio
			return c.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(response.Data{
			Message: "Login correcto",
			Data: fiber.Map{
				"token": authToken,
			},
		})
	}
}
