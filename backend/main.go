package main

import (
	"log"
	"main/users"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func main() {
	users.InitialMigration()
	app := fiber.New()

	app.Get("/login", Login)
	app.Post("/user", users.SaveUser)
	log.Fatal(app.Listen(":8000"))

}
