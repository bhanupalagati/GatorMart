package users

import (
	"gorm.io/gorm"
)

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
