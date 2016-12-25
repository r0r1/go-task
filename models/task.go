package models

import "github.com/jinzhu/gorm"

// Task Model
type Task struct {
	gorm.Model
	Name        string `json:"name" valid:"required"`
	UserID      uint   `gorm:"index" json:"user_id"`
	Description string `json:"description"`
	Parent      uint   `json:"parent_id"`
	Priority    int    `json:"priority"`
	Status      Status `json:"status_id"`
}
