package users

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Mysql@048@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	Title             string `json:"title"`
	SecondaryTitle    string `json:"secondarytitle"`
	Image             string `json:"image"`
	Price             string `json:"price"`
	SimpleDescription string `json:"simpleDescription"`
	Description       string `json:"description"`
	City              string `json:"city"`
	State             string `json:"state"`
	Location          string `json:"location"`
	Target            string `json:"target"`
	Category          string `json:"category"`
	PostedBy          string `json:"postedBy"`
	PostedDate        string `json:"postedDate"`
	Condition         string `json:"condition"`
	Age               string `json:"age"`
	AdStatus          string `json:"adStatus"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&User{})
}
