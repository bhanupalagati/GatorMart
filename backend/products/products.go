package products

import (
	"errors"
	"fmt"
	"log"
	"main/models"
	"os"
	"sort"
	"strconv"
	"strings"
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

	//swagger "github.com/arsmn/fiber-swagger/v2"
	_ "main/docs"
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
	// Name     string `json:"name"`
	// Email    string `json:"email"`
	// Password string `json:"password"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Profession string `json:"profession"`
	DOB        string `json:"DOB"`
}

// RegisterUser godoc
// @Summary Registers a user
// @Description Creates a GatorMart user account
// @Tags user
// @Accept  json
// @Produce  json
// @Param registerUser body RegisterRequest true "Register"
// @Success 200 {object} models.UserSwagger
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	registerUser := new(RegisterRequest)

	if err := c.BodyParser(registerUser); err != nil {
		fmt.Println("error")
		return c.Status(400).JSON(err.Error())
	}

	fmt.Println(registerUser.FirstName + "" + registerUser.Email + "-> " + string(registerUser.Password))

	var firstNameAlphaNumeric = true
	var lastNameAlphaNumeric = true
	for _, char := range registerUser.FirstName {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			firstNameAlphaNumeric = false
		}
	}
	for _, char := range registerUser.LastName {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			lastNameAlphaNumeric = false
		}
	}

	// check username pswdLength
	var firstNameLength bool
	if 5 <= len(registerUser.FirstName) && len(registerUser.FirstName) <= 50 {
		firstNameLength = true
	}
	var lastNameLength bool
	if 5 <= len(registerUser.LastName) && len(registerUser.LastName) <= 50 {
		lastNameLength = true
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
	if 8 < len(registerUser.Password) && len(registerUser.Password) < 60 {
		pswdLength = true
	}
	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nfirstNameAlphaNumeric:", firstNameAlphaNumeric, "\nlasttNameAlphaNumeric:", lastNameAlphaNumeric, "\nfirstNameLength:", firstNameLength, "\nlastNameLength:", lastNameLength)
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !firstNameAlphaNumeric || !lastNameAlphaNumeric || !firstNameLength || !lastNameLength {
		if !pswdLowercase {
			return c.Status(400).JSON(fiber.Map{
				"message": "password doesnt contain lowercase letter",
			})
		}
		if !pswdUppercase {
			return c.Status(400).JSON(fiber.Map{
				"message": "password doesnt contain uppercase letter",
			})
		}
		if !pswdNumber {
			return c.Status(400).JSON(fiber.Map{
				"message": "password doesnt contain any number",
			})
		}

		if !pswdSpecial {
			return c.Status(400).JSON(fiber.Map{
				"message": "password doesnt contain any special character",
			})
		}

		if !pswdLength {
			return c.Status(400).JSON(fiber.Map{
				"message": "password length should be greater than 8 and less than 60",
			})
		}

		if !pswdNoSpaces {
			return c.Status(400).JSON(fiber.Map{
				"message": "password shouldnot contains spaces",
			})
		}

		if !firstNameAlphaNumeric {
			return c.Status(400).JSON(fiber.Map{
				"message": "firstname should be alphanumeric",
			})
		}

		if !lastNameAlphaNumeric {
			return c.Status(400).JSON(fiber.Map{
				"message": "lastname should be alphanumeric",
			})
		}

		if !firstNameLength {
			return c.Status(400).JSON(fiber.Map{
				"message": "first name length should be greater than or equal to 5 and less than or equal to 50",
			})
		}

		if !lastNameLength {
			return c.Status(400).JSON(fiber.Map{
				"message": "last name length should be greater than or equal to 5 and less than or equal to 50",
			})
		}
		return c.Status(400).JSON(fiber.Map{
			"message": "please check username and password criteria",
		})
	}

	// dobDate, err := time.Parse("2006-01-02", registerUser.DOB)
	// if err != nil {
	// 	panic(err)
	// }
	s := strings.Split(registerUser.DOB, "-")
	year, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	month, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}
	date, err := strconv.Atoi(s[2])
	if err != nil {
		panic(err)
	}
	// timeFormat := "2006-01-02"
	dobDate := getDOB(year, month, date)

	password, _ := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	user := new(models.User)
	// user.UserID = int(user.ID)
	user.FirstName = registerUser.FirstName
	user.LastName = registerUser.LastName
	user.Email = registerUser.Email
	user.Password = password
	user.Profession = registerUser.Profession
	user.DOB = dobDate

	DB.Create(&user)

	return c.JSON(&user)

}
func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

// Login godoc
// @Summary Logins a user
// @Description Allows user to login to his account
// @Tags user
// @Accept  json
// @Produce  json
// @Param registerUser body RegisterRequest true "Register"
// @Success 200 {object} models.UserSwagger
// @Router /login [post]
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
	println(dbUser.ID)

	tokenExpirationTime := time.Now().Add(24 * time.Hour)
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["username"] = dbUser.FirstName
	claim["useremail"] = dbUser.Email
	claim["userid"] = dbUser.ID
	claim["ExpiresAt"] = tokenExpirationTime.Unix()
	claim["IssuedAt"] = time.Now().Unix()
	claim["Issuer"] = "GatorMart"
	claim["Subject"] = "Token for GatorMart frontend"
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer: strconv.Itoa(int(dbUser.ID)),
	// })

	token, err := claims.SignedString([]byte(KeyForAuthentication))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "user could not login",
		})
	}

	// cookie := fiber.Cookie{
	// 	Name:     "cookie",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }

	// c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
		"user":    dbUser,
		//"cookie":  cookie,
	})

}

// Logout godoc
// @Summary Logout a user
// @Description Allows user to logout of their account
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200
// @Router /logout [post]
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

// func UserDetails(c *fiber.Ctx) error {
// 	user, authorized := UserAuthorized(c)
// 	if !authorized {
// 		return c.Status(401).JSON("User not authorized")
// 	}
// 	return c.JSON(user)
// }

// func UserAuthorized(c *fiber.Ctx) (models.User, bool) {
// 	cookie := c.Cookies("cookie")

// 	authorised := false
// 	var user models.User

// 	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(KeyForAuthentication), nil

// 	})

// 	if err != nil {
// 		c.Status(fiber.StatusUnauthorized)
// 		// return c.JSON(fiber.Map{
// 		// 	"message": "User unauthenticated",
// 		// })
// 		return user, authorised
// 	}

// 	claims := token.Claims.(*jwt.StandardClaims)

// 	// var user models.User
// 	DB.Where("id=?", claims.Issuer).First(&user)
// 	if user.ID != 0 {
// 		authorised = true
// 	}

// 	return user, authorised
// }

func AuthorizeAndReturnUserDetails(c *fiber.Ctx) error {
	bearerToken := c.Get("authorization")
	println(bearerToken)
	tokenString := strings.Split(bearerToken, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KeyForAuthentication), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User unauthenticated",
		})
		println(token)
	}

	p := uint(claims["userid"].(float64))
	user := new(models.User)
	DB.Find(&user, p)
	return c.JSON(&user)
}

// SaveProduct godoc
// @Summary Saves a product post
// @Description Creates a product post in DB
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.ProductResponseSwagger true "SaveProduct"
// @Success 200 {object} models.ProductResponseSwagger
// @Router /product [post]
func SaveProduct(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
	bearerToken := c.Get("authorization")
	println(bearerToken)
	tokenString := strings.Split(bearerToken, " ")[1]
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	return KeyForAuthentication, nil
	// })
	// if err != nil{
	// 	return c.Status(400).JSON("Unable to parse token")
	// }
	// token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(KeyForAuthentication), nil

	// })
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KeyForAuthentication), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User unauthenticated",
		})
		println(token)
	}

	//claims := token.Claims.(*jwt.StandardClaims)

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	categoryTyp := CategoryType(product.Category)
	switch categoryTyp {
	case Automobile, Mobile, ElectronicsAppliances, Furniture, Books, Sports, Pets:

	}

	if err := IsValid(categoryTyp); err != nil {
		return c.Status(400).JSON("Invalid Category")
	}

	targetTyp := TargetType(product.Target)
	switch targetTyp {
	case Student, Farmer, Businessmen, Engineer, Other:

	}

	if err := IsValidTarget(targetTyp); err != nil {
		return c.Status(400).JSON("Invalid Target")
	}

	conditionTyp := ConditionType(product.Condition)
	switch conditionTyp {
	case Used, LikeNew, OpenBox, New:

	}

	if err := IsValidCondition(conditionTyp); err != nil {
		return c.Status(400).JSON("Invalid Condition")
	}

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	//product.PostedBy = claims["userid"].(string)
	p := uint(claims["userid"].(float64))
	println(p)
	product.PostedBy = p
	DB.Create(&product)
	return c.JSON(&product)
}

// GetCatergories godoc
// @Summary Get all categories
// @Description Returns all the product categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200
// @Router /categories [get]
func GetCatergories(c *fiber.Ctx) error {
	a := [...]string{"Automobile", "Mobile", "ElectronicsAppliances", "Furniture", "Books", "Sports", "Pets"}
	fmt.Println("a is ", a)
	return c.JSON(a)
}

// GetTarget godoc
// @Summary Get all target
// @Description Returns all the product target
// @Tags target
// @Accept  json
// @Produce  json
// @Success 200
// @Router /target [get]
func GetTarget(c *fiber.Ctx) error {
	a := [...]string{"Student", "Farmer", "Businessmen", "Engineer", "Other"}
	fmt.Println("a is ", a)
	return c.JSON(a)
}

// GetCondition godoc
// @Summary Get all conditions
// @Description Returns all the product conditions
// @Tags conditions
// @Accept  json
// @Produce  json
// @Success 200
// @Router /condition [get]
func GetCondition(c *fiber.Ctx) error {
	a := [...]string{"Used", "LikeNew", "OpenBox", "New"}
	fmt.Println("a is ", a)
	return c.JSON(a)
}

func IsValid(lt CategoryType) error {
	switch lt {
	case Automobile, Mobile, ElectronicsAppliances, Furniture, Books, Sports, Pets:
		return nil
	}
	return errors.New("Invalid Category type")
}

func IsValidTarget(lt TargetType) error {
	switch lt {
	case Student, Farmer, Businessmen, Engineer, Other:
		return nil
	}
	return errors.New("Invalid Target type")
}

func IsValidCondition(lt ConditionType) error {
	switch lt {
	case Used, LikeNew, OpenBox, New:
		return nil
	}
	return errors.New("Invalid Condition type")
}

// GetProducts godoc
// @Summary Get all products
// @Description Returns all the product posts
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductResponseSwagger
// @Router /products [get]
func GetProducts(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
	var products []models.Product
	DB.Find(&products)
	return c.JSON(&products)
}

// GetProductsByUser godoc
// @Summary get products posted by user
// @Description get products posted by user
// @Tags product
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductResponseSwagger
// @Router /productsByUser [get]
func GetProductsByUser(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
	bearerToken := c.Get("authorization")
	println(bearerToken)
	tokenString := strings.Split(bearerToken, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(KeyForAuthentication), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User unauthenticated",
		})
		println(token)
	}
	p := uint(claims["userid"].(float64))
	//	println(p)

	var products []models.Product
	DB.Where("posted_by = ?", p).Find(&products)
	//	println(len(products))
	return c.JSON(&products)
}

// GetProduct godoc
// @Summary get a product post by id
// @Description Get a product post by id
// @Tags product
// @Accept  json
// @Produce  json
// @Param id header int true "GetProduct"
// @Success 200 {object} models.ProductResponseSwagger
// @Router /product/:id [get]
func GetProduct(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
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

// UpdateProduct godoc
// @Summary update a product post
// @Description Update a product post in DB
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.ProductResponseSwagger true "UpdateProduct"
// @Success 200 {object} models.ProductResponseSwagger
// @Router /product/:id [put]
func UpdateProduct(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
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
	if product.Category != "" {
		categoryTyp := CategoryType(product.Category)
		switch categoryTyp {
		case Automobile, Mobile, ElectronicsAppliances, Furniture, Books, Sports, Pets:

		}

		if err := IsValid(categoryTyp); err != nil {
			return c.Status(400).JSON("Invalid Category")
		}
	}
	if product.Target != "" {
		targetTyp := TargetType(product.Target)
		switch targetTyp {
		case Student, Farmer, Businessmen, Engineer, Other:

		}

		if err := IsValidTarget(targetTyp); err != nil {
			return c.Status(400).JSON("Invalid Target")
		}
	}

	if product.Condition != "" {
		conditionTyp := ConditionType(product.Condition)
		switch conditionTyp {
		case Used, LikeNew, OpenBox, New:

		}

		if err := IsValidCondition(conditionTyp); err != nil {
			return c.Status(400).JSON("Invalid Condition")
		}
	}

	DB.Save(&product)
	return c.JSON(&product)

}

// DeleteProduct godoc
// @Summary delete a product post
// @Description delete a product post in DB
// @Tags product
// @Accept  json
// @Produce  json
// @Param id header int true "DeleteProduct"
// @Success 200
// @Router /product/:id [delete]
func DeleteProduct(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
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

// UploadImage godoc
// @Summary uploads image
// @Description Uploads image in Amazon S3
// @Tags product
// @Accept  json
// @Produce  json
// @Param files formData file true "UploadImage"
// @Success 200
// @Router /product/upload [post]
func UploadImage(c *fiber.Ctx) error {
	// _, authorized := UserAuthorized(c)
	// if !authorized {
	// 	return c.Status(401).JSON("User not authorized")
	// }
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

// FilterProducts godoc
// @Summary filters posts based on filtering criteria
// @Description filters posts based on filtering criteria
// @Tags product
// @Accept  json
// @Produce  json
// @Param filterConditions body models.Filter true "FilterProducts"
// @Success 200 {array} models.ProductResponseSwagger
// @Router /filterproducts [post]
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
	if filterConditions.SortBy != "" {
		if filterConditions.SortBy == "price" {
			sort.SliceStable(products, func(i, j int) bool {
				return products[i].Price < products[j].Price
			})
		}
		if filterConditions.SortBy == "date" {
			sort.SliceStable(products, func(i, j int) bool {
				// 	if products[i].CreatedAt.Day() != products[j].CreatedAt.Day(){
				// 	return products[i].CreatedAt.Day() < products[j].CreatedAt.Day()
				// 	}
				// if products[i].CreatedAt.Hour() != products[j].CreatedAt.Hour(){
				// 	return products[i].CreatedAt.Hour() < products[j].CreatedAt.Hour()
				// }
				return products[i].CreatedAt.Before(products[j].CreatedAt)
			})
		}
	}

	log.Println(fmt.Sprintf("Number of results obtained %d", len(products)))
	return c.JSON(&products)

}

type CategoryType string

type TargetType string

type ConditionType string

const (
	Automobile            CategoryType = "Automobile"
	Mobile                             = "Mobile"
	ElectronicsAppliances              = "ElectronicsAppliances"
	Furniture                          = "Furniture"
	Books                              = "Books"
	Sports                             = "Sports"
	Pets                               = "Pets"
)

const (
	Student     TargetType = "Student"
	Farmer                 = "Farmer"
	Businessmen            = "Businessmen"
	Engineer               = "Engineer"
	Other                  = "Other"
)

const (
	Used    ConditionType = "Used"
	LikeNew               = "LikeNew"
	OpenBox               = "OpenBox"
	New                   = "New"
)
