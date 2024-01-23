package main

import (
	"api.kyoris.com/database/mysql"
	"api.kyoris.com/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inicializamos Fiber
	app := fiber.New()

	// Inicializamos la base de datos
	db := mysql.NewMYSQLConnection("root:1234@tcp(127.0.0.1:3306)/NOMBRE_BASE_DE_DATOS_TEMPLATE?charset=utf8mb4&parseTime=True&loc=Local")

	// Hacemos las migraciones de la base de datos
	mysql.Migrate(db)

	// Inicializamos los servicios
	services := initServices(db)

	// Inicializamos las rutas
	routes.Setup(app, services)

	// Inicializamos el servidor
	app.Listen(":80")
}
