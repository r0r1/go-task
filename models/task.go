package models

import (
	"github.com/jinzhu/gorm"
)

// Task Model
type Task struct {
	gorm.Model
	Name        string `json:"name"`
	User        User   `gorm:"ForeignKey:UserID"`
	UserID      int    `json:"user_id"`
	Parent      int    `gorm:"default:0" json:"parent"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Tags        []Tag  `gorm:"many2many:tasks_tags;" json:"tags"`
}
