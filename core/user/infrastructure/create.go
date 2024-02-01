package userInfrastructure

import (
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
)

// Create implements userDomain.UserRepository.
func (r *userRepository) Create(user *userDomain.User) error {

	// Comprobamos si existe ya el email
	userModel := r.FindByEmail(user.Email)

	if userModel != nil {
		return userDomain.ErrEmailAlreadyExists
	}

	// Creamos la consulta para crear el usuario
	err := r.db.Model(&userDomain.User{}).Create(&user).Error

	if err != nil {
		return userDomain.ErrCreateUser
	}

	return nil
}
