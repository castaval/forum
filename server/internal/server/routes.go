package server

import "github.com/gofiber/fiber/v2"

func route(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "hello world!",
		})
	})
}
