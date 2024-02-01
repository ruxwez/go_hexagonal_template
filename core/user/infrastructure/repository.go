package userInfrastructure

import (
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db}
}
