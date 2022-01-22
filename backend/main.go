package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func main() {

	app := fiber.New()

	app.Get("/login", Login)

	log.Fatal(app.Listen(":8000"))

}
