package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"

	"github.com/nanafox/number-classifier-api-go/controllers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New(fiber.Config{AppName: "Number Classification API v1"})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, HEAD",
	}))

	api := app.Group("/api")
	api.Get("/classify-number", controllers.GetNumberClassification)

	log.Fatal(app.Listen(":" + port))
}
