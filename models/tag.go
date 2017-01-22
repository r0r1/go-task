package models

import "github.com/jinzhu/gorm"

// Tag Model
type Tag struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
}
