package router

import (
	"github.com/Omotolani98/github-insights/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes (app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Get("", handlers.Hello)
	v1.Get("/github/callback", handlers.GetUserAccessToken)
}
