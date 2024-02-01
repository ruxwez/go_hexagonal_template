package mysql

import (
	authDomain "github.com/ruxwez/go_hexagonal_template/core/auth/domain"
	userDomain "github.com/ruxwez/go_hexagonal_template/core/user/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.Table("users").AutoMigrate(&userDomain.User{})
	db.Table("sessions").AutoMigrate(&authDomain.Session{})
}
