package main

import (
	"forum/internal/server"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	server := server.New(app)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
