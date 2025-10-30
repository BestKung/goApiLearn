package config

import (
	"goApiLearn/controller"

	"github.com/gofiber/fiber/v2"
)

func InitApiConfig() {
	app := fiber.New()
	controller.NewCustomerController(app)
	app.Listen(":8080")
}
