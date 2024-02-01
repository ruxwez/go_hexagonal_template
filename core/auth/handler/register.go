package authHandler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	sharedDomain "github.com/ruxwez/go_hexagonal_template/core/_shared/domain"
	"github.com/ruxwez/go_hexagonal_template/core/_shared/domain/response"
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	emailUtils "github.com/ruxwez/go_hexagonal_template/utils/email"
	passwordUtils "github.com/ruxwez/go_hexagonal_template/utils/password"
)

// Get implements UserHandler.
func (a *authHandler) Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		bodyModel := authDomain.RegisterRequest{}

		// Parseamos el body a un modelo
		if err := ctx.BodyParser(&bodyModel); err != nil {
			// En el caso de error, devolvemos un error 400 y el codigo de error
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: sharedDomain.ErrBodyNotValid.Error(),
			})
		}

		// Eliminamos los espacion en blanco de los extremos
		bodyModel.Email = strings.TrimSpace(bodyModel.Email)
		bodyModel.Password = strings.TrimSpace(bodyModel.Password)
		bodyModel.Name = strings.TrimSpace(bodyModel.Name)

		// Comprobamos que el email sea valido
		if bodyModel.Email == "" {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmailNotValid.Error(),
			})
		}

		// Comprobamos que el email sea valido
		if !emailUtils.Check(bodyModel.Email) {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmailNotValid.Error(),
			})
		}

		// Comprobamos que la contraseña no este vacia, tenga al menos 8 caracteres y cumpla con los requisitos
		if bodyModel.Password == "" || len(bodyModel.Password) < 8 || !passwordUtils.CheckFormat(bodyModel.Password) {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrPasswordNotValid.Error(),
			})
		}

		// Comprobamos que el nombre no este vacio
		if bodyModel.Name == "" {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmptyName.Error(),
			})
		}

		// Creamos el modelo del usuario para hacer el Port de datos
		userModel := userDomain.User{
			Email:    bodyModel.Email,
			Password: bodyModel.Password,
			Name:     bodyModel.Name,
		}

		// Ejecutamos el servicio de registro de usuario
		token, err := a.authService.Register(&userModel, ctx.IP())

		// Comprobamos si el error es devido a que el email ya existe
		if errors.Is(err, userDomain.ErrEmailAlreadyExists) {
			return ctx.Status(http.StatusBadRequest).JSON(response.Error{
				Error: userDomain.ErrEmailAlreadyExists.Error(),
			})
		}

		// Comprobamos si ha sido otro error
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(response.Error{
				Error: sharedDomain.ErrInternalServerError.Error(),
			})
		}

		return ctx.Status(http.StatusOK).JSON(response.Data{
			Message: "Usuario registrado correctamente",
			Data:    fiber.Map{"token": token},
		})
	}
}
