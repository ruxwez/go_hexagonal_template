package main

import (
	sharedApplication "api.kyoris.com/core/_shared/application"
	sharedDomain "api.kyoris.com/core/_shared/domain"
	"gorm.io/gorm"
)

// Inicializamos los repsitorios para retornar un objeto de tipo Repositories
func initRepositories(dbConn *gorm.DB) *sharedDomain.Repositories {

	// Retornamos los repositorios inicializados para usarlos en los servicios
	return &sharedDomain.Repositories{}
}

// Inicializamos los servicios para retornar un objeto de tipo Services
func initServices(dbConn *gorm.DB) *sharedApplication.Services {
	// Obtenemos todos los Repositorios inicializados
	repositories := initRepositories(dbConn)

	// Retornamos los servicios inicializados para usarlos en los Handlers
	return &sharedApplication.Services{}
}
