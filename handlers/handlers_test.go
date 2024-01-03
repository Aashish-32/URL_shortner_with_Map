package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aashish-32/URL-ShortenerMap/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Dummy DataStructure
var (
	UrlStorage1 = make(map[string]string)
)

func TestShorteUrl(t *testing.T) {

	app := fiber.New()
	app.Post("/shorten", ShortenURL)

	payload := []byte(`{"long_url": "https://google.com"}`)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, fiber.StatusCreated, res.StatusCode)
	code := helpers.Base62Encode(len(UrlStorage1))

	expectedResponseBody := fmt.Sprintf(`{"long_url":"https://google.com","short_url":"http://localhost:8080/%s"}`, code)

	body, _ := io.ReadAll(res.Body)
	assert.Equal(t, expectedResponseBody, string(body))
}

func TestResolveURL(t *testing.T) {

	UrlStorage["abc"] = "https://example.com"

	app := fiber.New()
	app.Get("/:url", ResolveURL)

	req := httptest.NewRequest(http.MethodGet, "/abc", nil)
	res, err := app.Test(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, fiber.StatusFound, res.StatusCode)

	location := res.Header.Get("Location")
	assert.Equal(t, "https://example.com", location)
}
