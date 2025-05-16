package main

import (
	"github.com/Omotolani98/github-insights/config"
	"github.com/Omotolani98/github-insights/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main()  {
	appEnv := config.LoadEnv()

	app := fiber.New(fiber.Config{
		AppName: "Github Insights v1",
//		EnablePrintRoutes: true,
	})
	
	app.Use(cors.New())
	router.SetupRoutes(app)
	
	app.Listen(":" + appEnv.PORT)
}
