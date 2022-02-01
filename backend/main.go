package main

import (
	"log"
	"main/users"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func Routers(app *fiber.App) {
	app.Get("/users", users.GetUsers)
	app.Post("/user", users.SaveUser)
}

func main() {
	users.InitialMigration()
	app := fiber.New()
	Routers(app)
	// app.Get("/login", Login)
	// app.Post("/user", users.SaveUser)
	log.Fatal(app.Listen(":8000"))

}
