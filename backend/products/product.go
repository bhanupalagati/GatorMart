package products

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&product)
	return c.JSON(&product)
}
func GetProducts(c *fiber.Ctx) error {
	var products []Product
	DB.Find(&products)
	return c.JSON(&products)
}
