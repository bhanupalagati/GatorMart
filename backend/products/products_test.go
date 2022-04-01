package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

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

func TestDeleteProduct(t *testing.T) {
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

func TestLoginWhenPassWordInCorrect(t *testing.T) {
	var data = []byte(`{
		"email": "gatormart1@ufl.edu",
		"password": "gatormart1@A1"
	}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a wrong status code")
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
}

func TestLoginWhenPassWordCorrect(t *testing.T) {
	var data = []byte(`{
		"email": "gatormart1@ufl.edu",
		"password": "gatormart"
	}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a wrong status code")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestRegisterWhenSuccess(t *testing.T) {
	var data = []byte(`{
		"FirstName" : "Gowtham"
		"LastName": "Eda",
		"Email": "gatormart1@ufl.edu",
		"Password": "gatormart1@A1",
		"Profession": "Student",
		"DOB": "07211997"

	}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a sucess status code")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestRegisterWhenFailure(t *testing.T) {
	var data = []byte(`{
		"FirstName" : "Gowtham"
		"LastName": "",
		"Email": "gatormart1@ufl.edu",
		"Password": "gatormart1@A1",
		"Profession": "Student",
		"DOB": "07211997"

	}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a wrong status code")
	}

	assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
}

func TestLogoutWhenFailure(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/Logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a wrong status code")
	}

	assert.Equal(t, fiber.StatusInternalServerError, response.StatusCode)
}

func TestLogoutWhenSuccess(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/Logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a sucess status code")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestSaveProductWhenSucess(t *testing.T) {
	var data = []byte(`{
		"Category" = "ElectronicsAppliances"
}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/createGroup", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a sucess status code")
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
}

func TestSaveProductWhenFailure(t *testing.T) {
	var data = []byte(`{
		"Category" = "Bikes"
}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/createGroup", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a fail status code")
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
}

func TestGetProductsWhenSucess(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/Logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a sucess status code")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}

func TestGetProductsWhenFailure(t *testing.T) {
	var data = []byte(`{}`)

	app := fiber.New()

	req, _ := http.NewRequest("POST", "/api/Logout", bytes.NewBuffer(data))

	response, err := app.Test(req)

	if err != nil {
		t.Errorf("Handler Returned a fail status code")
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
}
