package server

import "github.com/gofiber/fiber/v2"

type server struct {
	app *fiber.App
}

func New(app *fiber.App) *server {
	return &server{
		app: app,
	}
}

func (s *server) Run() error {
	route(s.app)

	return s.app.Listen(":8080")
}
