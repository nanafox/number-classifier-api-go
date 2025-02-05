package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/nanafox/number-classifier-api-go/controllers"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Number Classification API v1"})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, HEAD",
	}))

	api := app.Group("/api")
	api.Get("/classify-number", controllers.GetNumberClassification)

	app.Listen(":3000")
}
