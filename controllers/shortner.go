package controllers

import (
	"fmt"

	"github.com/Aashish-32/URL-ShortenerMap/helpers"
	"github.com/gofiber/fiber/v2"

	"github.com/asaskevich/govalidator"
)

//A structure for the httpRequest payload

type RequestPayload struct {
	LongURL string `json:"long_url"`
}

var (
	UrlStorage = make(map[string]string)
)

func ShortenURL(c *fiber.Ctx) error {

	body := new(RequestPayload)

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot Parse JSON",
		})
	}
	if body.LongURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Empty long URL"})
	}
	if !govalidator.IsURL(body.LongURL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}

	shortURL := generateShortURL(body.LongURL)
	UrlStorage[shortURL] = body.LongURL
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"short_url": shortURL,
		"long_url":  body.LongURL,
	})

}

func ResolveURL(c *fiber.Ctx) error {
	shortcode := c.Params("url")

	longurl, exists := UrlStorage[shortcode]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found"})
	}

	return c.Redirect(longurl, fiber.StatusFound)
}

func generateShortURL(longURL string) string {

	id := len(UrlStorage)
	shortCode := helpers.Base62Encode(id)
	UrlStorage[shortCode] = longURL

	return fmt.Sprintf("http://localhost:8080/%s", shortCode)
}
