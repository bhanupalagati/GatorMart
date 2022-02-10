package main

import (
	"log"
	"main/products"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func Routers(app *fiber.App) {
	app.Get("/products", products.GetProducts)
	app.Post("/product", products.SaveProduct)
	app.Get("/product/:id", products.GetProduct)
	app.Put("/product/:id", products.UpdateProduct)
	app.Delete("/product/:id", products.DeleteProduct)
}

func main() {
	products.InitialMigration()
	app := fiber.New()
	app.Use(cors.New())
	Routers(app)
	// app.Get("/login", Login)
	// app.Post("/user", users.SaveUser)
	log.Fatal(app.Listen(":8000"))

}
