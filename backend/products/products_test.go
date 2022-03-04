package products

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSaveProduct(t *testing.T) {
	tests := []struct {
		description string // description of the test case

		route string // route path to test

		expectedCode int // expected HTTP status code

	}{

		// First test case

		{

			description: "get HTTP status 200",

			route: "/product",

			expectedCode: 200,
		},

		// Second test case

		{

			description: "get HTTP status 404, when route is not exists",

			route: "/not-found",

			expectedCode: 404,
		},
	}

	// Define Fiber app.

	app := fiber.New()

	//product := new(models.Product)

	// Create route with GET method for test

	app.Get("/product", func(c *fiber.Ctx) error {

		// Return simple string as response

		return c.SendStatus(200)

	})

	// Iterate through test single test cases

	for _, test := range tests {

		// Create a new http request with the route from the test case

		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,

		// the second argument is a request latency

		// (set to -1 for no latency)

		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}

func TestGetProducts(t *testing.T) {
	tests := []struct {
		description string // description of the test case

		route string // route path to test

		expectedCode int // expected HTTP status code

	}{

		// First test case

		{

			description: "get HTTP status 200",

			route: "/product",

			expectedCode: 200,
		},
	}

	// Define Fiber app.

	app := fiber.New()

	//product := new(models.Product)

	// Create route with GET method for test

	app.Get("/product", func(c *fiber.Ctx) error {

		// Return simple string as response

		return c.SendStatus(200)

	})

	// Iterate through test single test cases

	for _, test := range tests {

		// Create a new http request with the route from the test case

		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,

		// the second argument is a request latency

		// (set to -1 for no latency)

		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}

func TestGetProduct(t *testing.T) {
	tests := []struct {
		description string // description of the test case

		route string // route path to test

		expectedCode int // expected HTTP status code

	}{

		// First test case

		{

			description: "get HTTP status 200",

			route: "/product",

			expectedCode: 200,
		},
	}

	// Define Fiber app.

	app := fiber.New()

	//product := new(models.Product)

	// Create route with GET method for test

	app.Get("/product", func(c *fiber.Ctx) error {

		// Return simple string as response

		return c.SendStatus(200)

	})

	// Iterate through test single test cases

	for _, test := range tests {

		// Create a new http request with the route from the test case

		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,

		// the second argument is a request latency

		// (set to -1 for no latency)

		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}

}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		description string // description of the test case

		route string // route path to test

		expectedCode int // expected HTTP status code

	}{

		// First test case

		{

			description: "get HTTP status 200",

			route: "/product",

			expectedCode: 200,
		},
	}

	// Define Fiber app.

	app := fiber.New()

	//product := new(models.Product)

	// Create route with GET method for test

	app.Get("/product", func(c *fiber.Ctx) error {

		// Return simple string as response

		return c.SendStatus(200)

	})

	// Iterate through test single test cases

	for _, test := range tests {

		// Create a new http request with the route from the test case

		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,

		// the second argument is a request latency

		// (set to -1 for no latency)

		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}
