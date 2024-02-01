package authDomain

import "errors"

var (
	// Codigo de error para cuando falla la creación de una sesión
	ErrSessionCreation = errors.New("session_creation_error")

	// Codigo de error para cuando falta el token
	ErrEmptyToken = errors.New("empty_token")

	// Codigo de error para cuando el token no es válido
	ErrInvalidToken = errors.New("invalid_token")

	// Codigo de error para cuando no se ha podido eliminar la sesión
	ErrSessionDeletion = errors.New("session_deletion_error")
)
