package tokenUtils

import "strings"

func ValidToken(token string) (string, bool) {

	// Limpiamos los espacios en blanco sobrantes
	token = strings.TrimSpace(token)

	// Validamos que tenga método "Bearer "
	if !strings.Contains(token, "Bearer ") {
		return "", false
	}

	// Quitamos el Bearer de la cadena de texto
	token = strings.Replace(token, "Bearer ", "", 1)

	return token, true
}
