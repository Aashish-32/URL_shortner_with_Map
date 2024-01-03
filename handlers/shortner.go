package handlers

import (
	"fmt"

	"github.com/Aashish-32/URL-ShortenerMap/helpers"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

// A structure for the HTTP request payload
type RequestPayload struct {
	LongURL string `json:"long_url"`
}

// Data structure for in-memory storage of URLs
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
	id := len(UrlStorage)
	shortCode := helpers.Base62Encode(id)

	//checking for repeated URLS

	for shortcode, Storedlongurl := range UrlStorage {
		if Storedlongurl == longURL {
			return fmt.Sprintf("http://localhost:8080/%s", shortcode)
		}
	}

	//appending htttp so that user can shorten urls without having to type http:// or https://

	if len(longURL) >= 8 && longURL[:7] != "http://" && longURL[:8] != "https://" {
		longURL = "http://" + longURL
	}

	UrlStorage[shortCode] = longURL

	return fmt.Sprintf("http://localhost:8080/%s", shortCode)
}
