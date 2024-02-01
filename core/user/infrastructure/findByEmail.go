package userInfrastructure

import (
	"errors"

	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	"gorm.io/gorm"
)

// FindByEmail implements userDomain.UserRepository.
func (r *userRepository) FindByEmail(email string) *userDomain.User {
	user := userDomain.User{}

	err := r.db.Model(&user).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
