package main

import (
	"github.com/Aashish-32/URL-ShortenerMap/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.SetupRoutes(app)

	panic(app.Listen(":8080"))
}
