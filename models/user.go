package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `valid:"Required"`
	Email    string `gorm:"type:varchar(100);unique_index" valid:"Required;Email"`
	Username string `gorm:"type:varchar(100);unique_index" valid:"Required"`
	Password string `valid:"Required"`
	Contact  string
	Address  string `gorm:"type:text"`
	Task     []Task `gorm:"ForeignKey:UserId"`
}
