package main

import (
	"log"

	"github.com/Kalveir/go-fiber-api/database"
	"github.com/gofiber/fiber/v3"
	// "github.com/Kalveir/go-fiber-api/config"
	"github.com/Kalveir/go-fiber-api/handler"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: true,
		AppName:      "test",
	})

	app.Use(cors.New())

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

	//routes
	app.Get("/book", handler.GetBooks)
	app.Get("/book/:id", handler.FindBooks)
	app.Post("/book/store", handler.StoreBooks)
	app.Put("/book/:id", handler.UpdateBooks)
	app.Delete("/book/:id", handler.DeleteBooks)
}
