package sharedDomain

import "errors"

var (

	// Problema interno del servidor
	ErrInternalServerError = errors.New("internal_server_error")

	// Codigo de error cuando el body no es correcto
	ErrBodyNotValid = errors.New("body_not_valid")
)
