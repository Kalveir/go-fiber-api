package main

import (
	"log"

	"github.com/Kalveir/go-fiber-api/database"
	"github.com/gofiber/fiber/v3"
	"github.com/Kalveir/go-fiber-api/route"
	// "github.com/Kalveir/go-fiber-api/config"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	database.ConnectDb()
	// app := fiber.New(fiber.Config{fiberConfig()})
	app := fiber.New()

	app.Use(cors.New())

	route.setupRoutes(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

}
