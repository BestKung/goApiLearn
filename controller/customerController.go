package controller

import (
	"goApiLearn/model"

	"github.com/gofiber/fiber/v2"
)

func GetCustomer(app *fiber.App) {
	app.Get("/customers", func(c *fiber.Ctx) error {
		custs := []model.Customer{
			{"Bii", "18", "000999384", "22/xxx", "CUS625"},
		}
		return c.JSON(custs)
	})
}
