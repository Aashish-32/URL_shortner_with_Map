package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	shortcode := c.Params("url")

	longurl, exists := UrlStorage[shortcode]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found"})
	}

	return c.Redirect(longurl, fiber.StatusFound)
}
