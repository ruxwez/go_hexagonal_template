package passwordUtils

import "golang.org/x/crypto/bcrypt"

func Validate(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
