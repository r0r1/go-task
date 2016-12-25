package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Status Model
type Status struct {
	gorm.Model
	Name      string `json:"name" valid:"required"`
	Label     string `json:"label"`
	CreatedAt time.Time
}
