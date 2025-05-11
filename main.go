package main

import (
	"github.com/Omotolani98/github-insights/config"
	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func main()  {
	appEnv := config.LoadEnv()

	app := fiber.New(fiber.Config{
		AppName: "Github Insights v1",
		EnablePrintRoutes: true,
	})

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Server Healthy")
	})

	log.Infof("Insight Server is running at <::> %s", appEnv.PORT)
	app.Listen(":" + appEnv.PORT)
}
