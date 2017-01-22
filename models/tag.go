package models

// Tag Model
type Tag struct {
	ID   uint
	Name string `json:"name" binding:"required"`
}
