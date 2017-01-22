package models

import (
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique_index" json:"email"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}
