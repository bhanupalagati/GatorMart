package products

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Mysql@048@tcp(127.0.0.1:3306)/godbgator1?charset=utf8mb4&parseTime=True&loc=Local"

type Product struct {
	gorm.Model
	Title             string `json:"title"`
	SecondaryTitle    string `json:"secondary_title"`
	ImageUrl          string `json:"imageUrl"`
	Price             int    `json:"price"`
	SimpleDescription string `json:"simple_desc"`
	Description       string `json:"description"`
	City              string `json:"city"`
	State             string `json:"state"`
	Location_Lat      string `json:"location_lat"`
	Location_Long     string `json:"location_long"`
	Target            string `json:"target"`
	Category          string `json:"category"`
	PostedBy          string `json:"posted_by"`
	PostedDate        string `json:"posted_date"`
	Condition         string `json:"condition"`
	Age               int    `json:"age"`
	AdStatus          string `json:"status"`
	Images            string `json:"images"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Product{})
}

func SaveProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	DB.Create(&product)
	return c.JSON(&product)
}
func GetProducts(c *fiber.Ctx) error {
	var products []Product
	DB.Find(&products)
	return c.JSON(&products)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	var product Product
	DB.Find(&product, id)
	if product.ID == 0 {
		return errors.New("product id doesnt exist")
	}
	return c.JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	product := new(Product)
	DB.First(&product, id)
	if product.ID == 0 {
		return errors.New("product id doesnt exist")
	}
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	DB.Save(&product)
	return c.JSON(&product)

}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	var product Product
	DB.First(&product, id)
	if product.ID == 0 {
		return errors.New("product id doesnt exist")
	}
	DB.Delete(&product)
	return c.SendString("Product ad is deleted")
}

func HandleFileUpload(c *fiber.Ctx) error {

	// parse incomming image file

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:8000/images/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})
}
