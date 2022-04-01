package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID     int    `gorm:"primary_key" json:"userId"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   []byte `json:"-"`
	Profession string `json:"profession"`
	DOB        time.Time
}
