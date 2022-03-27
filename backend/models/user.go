package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   []byte `json:"-"`
	Profession string `json:"profession"`
	DOB        time.Time
}
