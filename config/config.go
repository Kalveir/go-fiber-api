package config

import (
	"github.com/gofiber/fiber/v3"
)

func fiberConfig() fiber.Config {
	return fiber.Config {
		AppName : "anjay",
		// Prefork: true,
	}
}