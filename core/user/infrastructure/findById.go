package userInfrastructure

import (
	"errors"

	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	"gorm.io/gorm"
)

// FindByEmail implements userDomain.UserRepository.
func (r *userRepository) FindById(id uint) *userDomain.User {
	user := userDomain.User{}

	err := r.db.Model(&user).Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &user
}
