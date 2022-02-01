package main

import (
	"log"
	"main/products"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func Routers(app *fiber.App) {
	app.Get("/products", products.GetProducts)
	app.Post("/product", products.SaveProduct)
}

func main() {
	products.InitialMigration()
	app := fiber.New()
	Routers(app)
	// app.Get("/login", Login)
	// app.Post("/user", users.SaveUser)
	log.Fatal(app.Listen(":8000"))

}
