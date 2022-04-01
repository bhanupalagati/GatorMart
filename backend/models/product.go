package models

import (
	"gorm.io/gorm"
)

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
	PostedBy          uint   `json:"posted_by"`
	PostedDate        string `json:"posted_date"`
	Condition         string `json:"condition"`
	Age               int    `json:"age"`
	AdStatus          string `json:"status"`
	Images            string `json:"images"`
}
