package emailUtils

import "regexp"

// Check verifica si una dirección de correo electrónico es válida.
// Utiliza una expresión regular para verificar la sintaxis del correo electrónico
// basada en la especificación del formato de correo electrónico (RFC 5322).
func Check(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
