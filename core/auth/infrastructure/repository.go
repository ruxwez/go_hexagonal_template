package authInfrastructure

import (
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) authDomain.AuthRepository {
	return &authRepository{db}
}
