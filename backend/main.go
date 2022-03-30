package main

import (
	"log"
	"main/products"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

const KeyForAuthentication = "secret"

func Login(c *fiber.Ctx) error {

	return c.SendString("hi")
}

func Routers(app *fiber.App) {
	app.Post("/register", products.Register)
	app.Post("/login", products.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	// app.Get("/products", products.GetProducts)
	app.Post("/product", products.SaveProduct)
	app.Get("/product/:id", products.GetProduct)
	app.Put("/product/:id", products.UpdateProduct)
	app.Delete("/product/:id", products.DeleteProduct)
	app.Post("/product/upload", products.UploadImage)

	app.Post("/logout", products.Logout)
	app.Post("/filterproducts", products.FilterProducts)
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	app.Get("/products", products.GetProducts)
	app.Get("/categories", products.GetCatergories)
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: []byte("secret"),
	// }))
	// app.Post("/product", products.SaveProduct)
}

// func authReq() func(c *fiber.Ctx){
// 	return jwtware.New(jwtware.Config{
// 		SigningKey: []byte(KeyForAuthentication),
// 	})
// }

func main() {
	products.InitialMigration()
	app := fiber.New()
	app.Use(cors.New())
	//app.Use(middleware.Logger())

	Routers(app)

	// awsAccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	// fmt.Println("My access key ID is ", awsAccessKeyID)

	// app.Get("/login", Login)
	// app.Post("/user", users.SaveUser)
	log.Fatal(app.Listen(":8000"))

}
