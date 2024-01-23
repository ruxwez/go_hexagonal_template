package passwordUtils

import "unicode"

func CheckFormat(pass string) bool {
	var mayus, minus, number, special bool

	for _, char := range pass {
		if unicode.IsUpper(char) {
			mayus = true
			continue
		}

		if unicode.IsLower(char) {
			minus = true
			continue
		}

		if unicode.IsNumber(char) {
			number = true
			continue
		}

		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			special = true
			continue
		}
	}

	return mayus && minus && number && special
}
