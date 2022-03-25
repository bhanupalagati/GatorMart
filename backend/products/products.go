package products

import (
	"errors"
	"fmt"
	"log"
	"main/models"
	"os"
	"strconv"
	"time"
	"unicode"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.User{})
}

//User Management
const KeyForAuthentication = "secret"

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	registerUser := new(RegisterRequest)

	if err := c.BodyParser(registerUser); err != nil {
		fmt.Println("error")
		return c.Status(400).JSON(err.Error())
	}

	fmt.Println(registerUser.Name + "" + registerUser.Email + "-> " + string(registerUser.Password))

	var nameAlphaNumeric = true
	for _, char := range registerUser.Name {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
	}

	// check username pswdLength
	var nameLength bool
	if 5 <= len(registerUser.Name) && len(registerUser.Name) <= 50 {
		nameLength = true
	}
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range string(registerUser.Password) {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	fmt.Println(len(registerUser.Password))
	if 11 < len(registerUser.Password) && len(registerUser.Password) < 60 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {

		return c.Status(400).JSON(fiber.Map{
			"message": "please check username and password criteria",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	user := new(models.User)
	user.Name = registerUser.Name
	user.Email = registerUser.Email
	user.Password = password

	DB.Create(&user)

	return c.JSON(&user)

}

func Login(c *fiber.Ctx) error {
	user := new(RegisterRequest)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	//var user models.User
	dbUser := new(models.User)
	fmt.Println(user.Email + "->")
	DB.Where("email= ?", user.Email).First(&dbUser)
	//DB.First(&dbUser, usr.Email)
	//database.DB.Where("email= ?", data["email"]).First(&user)
	fmt.Println(dbUser.ID, dbUser.Email)
	if dbUser.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	fmt.Println("====")
	fmt.Println(dbUser.Password)
	fmt.Println("====")
	fmt.Println(user.Password)
	if err := bcrypt.CompareHashAndPassword(dbUser.Password, []byte(user.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(dbUser.ID)),
	})

	token, err := claims.SignedString([]byte(KeyForAuthentication))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "user could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "cookie",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
		"cookie":  cookie,
	})

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "cookie",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout successful ",

		"cookie": cookie,
	})
}

func UserAuthorized(c *fiber.Ctx) (models.User, bool) {
	cookie := c.Cookies("cookie")

	authorised := false
	var user models.User

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(KeyForAuthentication), nil

	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		// return c.JSON(fiber.Map{
		// 	"message": "User unauthenticated",
		// })
		return user, authorised
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// var user models.User
	DB.Where("id=?", claims.Issuer).First(&user)
	if user.ID != 0 {
		authorised = true
	}

	return user, authorised
}

func SaveProduct(c *fiber.Ctx) error {
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	DB.Create(&product)
	return c.JSON(&product)
}
func GetProducts(c *fiber.Ctx) error {
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
	var products []models.Product
	DB.Find(&products)
	return c.JSON(&products)
}

func GetProduct(c *fiber.Ctx) error {
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	var product models.Product
	DB.Find(&product, id)
	if product.ID == 0 {
		return errors.New("product id doesnt exist")
	}
	return c.JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	product := new(models.Product)
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
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is an integer")
	}
	var product models.Product
	DB.First(&product, id)
	if product.ID == 0 {
		return errors.New("product id doesnt exist")
	}
	DB.Delete(&product)
	return c.JSON("Product ad is deleted")
}

func UploadImage(c *fiber.Ctx) error {
	_, authorized := UserAuthorized(c)
	if !authorized {
		return c.Status(401).JSON("User not authorized")
	}
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

func FilterProducts(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
	filterConditions := new(models.Filter)
	if err := c.BodyParser(filterConditions); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var products []models.Product
	res := DB.Find(&products)
	if filterConditions.Title != "" && filterConditions != nil {
		res.Where("title LIKE ?", fmt.Sprintf("%%%s%%", filterConditions.Title)).Find(&products)
	}
	if filterConditions.Target != "" && filterConditions != nil {
		res.Where("target = ?", filterConditions.Target).Find((&products))
	}
	if filterConditions.Category != "" && filterConditions != nil {
		res.Where("category = ?", filterConditions.Category).Find((&products))
	}
	price := filterConditions.Price

	if price.Operator != "" {
		if price.Operator == "=" {
			res.Where("price = ?", filterConditions.Price.FromValue).Find((&products))

		} else if price.Operator == ">" {
			res.Where("price > ?", filterConditions.Price.FromValue).Find((&products))

		} else if price.Operator == "<" {
			res.Where("price < ?", filterConditions.Price.ToValue).Find((&products))
		} else if price.Operator == "between" {
			res.Where(" price > ?", filterConditions.Price.FromValue).Where("price < ?", filterConditions.Price.ToValue).Find((&products))
		}

	}

	log.Println(fmt.Sprintf("Number of results obtained %d", len(products)))
	return c.JSON(&products)

}
