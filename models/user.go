package models

import "github.com/jinzhu/gorm"

// User Model
type User struct {
	gorm.Model
	Name     string `json:"name" valid:"required"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
	Task     []Task `gorm:"ForeignKey:UserId"`
}
