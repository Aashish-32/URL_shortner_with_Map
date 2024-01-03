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

// Data structure for in-memory storage o URLs

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

func generateShortURL(longURL string) string {

	//Checking for Repeated URL

	for shortcode, longcode := range UrlStorage {
		if longcode == longURL {
			return fmt.Sprintf("http://localhost:8080/%s", shortcode)

		}
	}

	id := len(UrlStorage)
	shortCode := helpers.Base62Encode(id)

	UrlStorage[shortCode] = longURL

	return fmt.Sprintf("http://localhost:8080/%s", shortCode)
}
