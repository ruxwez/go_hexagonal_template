package main

import (
	sharedApplication "github.com/ruxwez/go_hexagonal_template/core/_shared/application"
	sharedDomain "github.com/ruxwez/go_hexagonal_template/core/_shared/domain"
	authApplication "github.com/ruxwez/go_hexagonal_template/core/auth/application"
	authInfrastructure "github.com/ruxwez/go_hexagonal_template/core/auth/infrastructure"

	userApplication "github.com/ruxwez/go_hexagonal_template/core/user/application"
	userInfrastructure "github.com/ruxwez/go_hexagonal_template/core/user/infrastructure"
	"gorm.io/gorm"
)

// Inicializamos los repsitorios para retornar un objeto de tipo Repositories
func initRepositories(dbConn *gorm.DB) *sharedDomain.Repositories {
	userRepository := userInfrastructure.NewMySQLRepository(dbConn)
	authRepository := authInfrastructure.NewMySQLRepository(dbConn)

	// Retornamos los repositorios inicializados para usarlos en los servicios
	return &sharedDomain.Repositories{
		UserRepository: userRepository,
		AuthRepository: authRepository,
	}
}

// Inicializamos los servicios para retornar un objeto de tipo Services
func initServices(dbConn *gorm.DB) *sharedApplication.Services {
	// Obtenemos todos los Repositorios inicializados
	repositories := initRepositories(dbConn)

	userService := userApplication.NewService(repositories)
	authService := authApplication.NewAuthService(repositories)

	// Retornamos los servicios inicializados para usarlos en los Handlers
	return &sharedApplication.Services{
		UserService: userService,
		AuthService: authService,
	}
}
