package userDomain

import "errors"

var ()

// Todos los errores relacionados con el usuario
var (
	// Codigo de error para cuando el usuario no esta activo
	ErrUserNotActive = errors.New("user_not_active")

	// Codigo de error para cuando el usuario no existe
	ErrNotExistsUser = errors.New("not_exists_user")

	// Codigo de error para cuando no se puede crear el usuario
	ErrCreateUser = errors.New("cant_create_user")

	// Codigo de error para cuando las credenciales son incorrectas
	ErrIncorrectCredentials = errors.New("incorrect_credentials")

	// Mensaje de error para cuando el usuario ya existe.
	ErrEmailAlreadyExists = errors.New("email_already_exists")

	// Mensaje de error para cuando el nombre de usuario ya existe.
	ErrUsernameAlreadyExists = errors.New("username_already_exists")

	// Mensaje de error para cuando no se ha podido obtener los datos del usuario
	ErrGetUserData = errors.New("error_get_user_data")
)

// Todos los errores relacionados con los campos del usuario
var (
	// Codigo de error para cuando el email no es valido
	ErrEmailNotValid = errors.New("email_not_valid")

	// Codigo de error para cuando el email esta vacio
	ErrEmptyEmail = errors.New("empty_email")

	// Codigo de error para cuando el nombre de usuario esta vacio
	ErrEmptyUsername = errors.New("empty_username")

	// Codigo de error para cuando el nombre esta vacio
	ErrEmptyName = errors.New("empty_firstname")

	// Codigo de error para cuando el genero no es valido
	ErrIncorrectGender = errors.New("incorrect_gender")

	// Codigo de error para cuando la contraseña es incorrecta
	ErrIncorrectPassword = errors.New("incorrect_password")

	// Codigo de error para cuando la contraseña no cumple con los requisitos
	ErrPasswordNotValid = errors.New("password_not_valid")

	// Codigo de error para cuando falle la encriptacion de la contraseña
	ErrEncryptPassword = errors.New("encrypt_password")

	// Codigo de error para cueando la fecha no es valida
	ErrBirthdayNotValid = errors.New("birthday_not_valid")
)
