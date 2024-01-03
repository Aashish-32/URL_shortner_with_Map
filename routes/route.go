package routes

import (
	"github.com/Aashish-32/URL-ShortenerMap/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/shorten", handlers.ShortenURL)
	app.Get("/:url", handlers.ResolveURL)

}
