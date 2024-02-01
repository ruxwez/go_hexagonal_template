package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ruxwez/go_hexagonal_template/database/mysql"
	"github.com/ruxwez/go_hexagonal_template/routes"
)

func main() {
	// Inicializamos Fiber
	app := fiber.New()

	// Inicializamos la base de datos
	db := mysql.NewMYSQLConnection("root:1234@tcp(127.0.0.1:3306)/example?charset=utf8mb4&parseTime=True&loc=Local")

	// Hacemos las migraciones de la base de datos
	mysql.Migrate(db)

	// Inicializamos los servicios
	services := initServices(db)

	// Inicializamos las rutas
	routes.Setup(app, services)

	// Inicializamos el servidor
	app.Listen(":80")
}
