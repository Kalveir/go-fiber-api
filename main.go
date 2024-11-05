package main

import (
	"log"

	"github.com/Kalveir/go-fiber-api/config"
	"github.com/Kalveir/go-fiber-api/database"
	"github.com/Kalveir/go-fiber-api/route"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnectDb()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareSetup(app)

	route.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

}
