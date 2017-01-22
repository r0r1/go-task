package models

import (
	"github.com/jinzhu/gorm"
)

// Note Model
type Note struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Task        Task   `gorm:"ForeignKey:TaskID"`
	TaskID      int    `json:"task_id"`
	User        User   `gorm:"ForeignKey:UserID"`
	UserID      int    `json:"user_id"`
}
