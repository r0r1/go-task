package models

import (
	"github.com/jinzhu/gorm"
)

// Task Model
type Task struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	User        User   `gorm:"ForeignKey:UserID"`
	UserID      int    `json:"user_id"`
	Parent      int    `gorm:"default:0" json:"parent"`
	Priority    int    `json:"priority" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Description string `json:"description"`
}
