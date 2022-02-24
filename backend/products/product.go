package products

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var AccessKeyID string
var SecretAccessKey string
var MyRegion string
var MyBucket string
var filepath string

func ConnectAws() *session.Session {
	AccessKeyID = GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion = GetEnvWithKey("AWS_REGION")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(MyRegion),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

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
		return c.Status(400).JSON(err.Error())
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
		return c.Status(400).JSON(err.Error())
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
	return c.JSON("Product ad is deleted")
}

func UploadImage(c *fiber.Ctx) error {
	LoadEnv()
	sess := ConnectAws()
	uploader := s3manager.NewUploader(sess)
	MyBucket = GetEnvWithKey("BUCKET_NAME")
	form, err := c.MultipartForm()
	if err != nil {

		log.Println("error while reading mutipart form, --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}
	files := form.File["photo"]

	var pathUrls []string
	var imageNameList []string
	//var headerList []string
	var sizeList []int64

	for _, file := range files {

		filename := file.Filename
		//fmt.Println(" Reached 2 filename :", filename)
		f, err := file.Open()
		if err != nil {

			log.Println("error while trying to Open file --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
		}
		//filename :=file.Header.Filename
		//upload to the s3 bucket
		up, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(MyBucket),
			ACL:    aws.String("private"),
			Key:    aws.String(filename),
			Body:   f,
		})
		if err != nil {

			log.Println("error while uploading image to S3 bucket --> ", err)
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "uploader": up, "data": nil})
		}
		//filepath = "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
		filepath = "https://" + MyBucket + "." + "s3" + ".amazonaws.com/" + filename

		pathUrls = append(pathUrls, filepath)
		imageNameList = append(imageNameList, filename)
		sizeList = append(sizeList, file.Size)

	}

	//file, header, err := c.FormFile("photo")
	// file, err := c.FormFile("photo")
	// fmt.Println(" Reached 1")
	// if err != nil {

	// 	log.Println("image save error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }
	// fmt.Println(" Reached 2-0")
	//filename := file.Filename
	// fmt.Println(" Reached 2 filename :", filename)
	// f, err := file.Open()
	// if err != nil {

	// 	log.Println("image save error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// }
	// //filename :=file.Header.Filename
	// //upload to the s3 bucket
	// up, err := uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(MyBucket),
	// 	ACL:    aws.String("public-read"),
	// 	Key:    aws.String(filename),
	// 	Body:   f,
	// })
	// if err != nil {

	// 	log.Println("tttimage save error --> ", err)
	// 	return c.JSON(fiber.Map{"status": 500, "message": "Server error", "uploader": up, "data": nil})
	// }
	// filepath = "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
	// //  c.JSON(http.StatusOK, gin.H{
	// //   "filepath":    filepath,
	// //  })
	data := map[string]interface{}{

		"imageName": imageNameList,
		"imageUrl":  pathUrls,
		// "header":   file.Header,
		"size": sizeList,
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data": data})

}
