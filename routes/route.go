package routes

import (
	"github.com/Aashish-32/URL-ShortenerMap/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", controllers.ShortenURL)
	app.Get("/:url", controllers.ResolveURL)

}
