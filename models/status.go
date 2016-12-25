package models

import "github.com/jinzhu/gorm"

// Status Model
type Status struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Label string `json:"label"`
}
