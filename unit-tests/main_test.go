package main

import (
	"net/http"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()

	app.Get("/v1", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	return app
}

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description string
		route string
		expectedCode int
	}{
		{
			description: "index route test",
			route: "/v1",
			expectedCode: 200,
		},
	}

	log.Infof("Length :: %d", len(tests))
	app := setupApp()

	for _, test := range tests {
		req, _ := http.NewRequest("GET", test.route, nil)

		_, _ := app.Test(req, -1)
		assert.Equalf(t, test.expectedCode, test.description, test.description)
	}
}
