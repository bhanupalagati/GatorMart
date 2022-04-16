package main

import (
	"log"
	"main/products"

	//swagger "github.com/arsmn/fiber-swagger"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	//swagger "github.com/arsmn/fiber-swagger/v2"
	_ "main/docs"

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
	app.Get("/target", products.GetTarget)
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

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api
// @schemes http
func main() {
	products.InitialMigration()
	app := fiber.New()
	app.Use(cors.New())
	//app.Use(middleware.Logger())
	app.Get("/swagger/*", swagger.HandlerDefault)
	Routers(app)

	// awsAccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	// fmt.Println("My access key ID is ", awsAccessKeyID)

	// app.Get("/login", Login)
	// app.Post("/user", users.SaveUser)
	// default
	log.Fatal(app.Listen(":8000"))

}
