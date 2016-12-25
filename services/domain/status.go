package domain

import (
	"fmt"
	"go-task/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Status Storage
type StatusStorage struct {
	db *gorm.DB
}

// New Status Storage
func NewStatusStorage(db *gorm.DB) *StatusStorage {
	return &StatusStorage{db}
}

// GetAll of the chocolate
func (s StatusStorage) GetAll() []models.Status {
	var statuses []models.Status
	s.db.Order("id").Find(&statuses)
	return statuses
}

// GetOne tasty chocolate
func (s StatusStorage) GetOne(id string) (models.Status, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.Status{}, fmt.Errorf("Chocolate id must be integer: %s", id)
	}
	var status models.Status
	s.db.First(&status, intID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return models.Status{}, fmt.Errorf("Chocolate for id %s not found", id)
	}
	return status, nil
}

// Insert a fresh one
func (s *StatusStorage) Insert(status models.Status) uint {
	s.db.Create(&status)
	return status.ID
}

// Delete one :(
func (s *StatusStorage) Delete(id string) error {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return fmt.Errorf("Chocolate id must be integer: %s", id)
	}

	var status models.Status
	s.db.First(&status, intID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Chocolate with id %s does not exist", id)
	}
	s.db.Delete(&status)

	return s.db.Error
}

// Update updates an existing chocolate
func (s *StatusStorage) Update(c models.Status) error {
	var status models.Status
	s.db.First(&status, c.ID)
	if err := s.db.Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("Chocolate with id %s does not exist", c.ID)
	} else if err != nil {
		return err
	}
	status.Name = c.Name
	status.Label = c.Label
	s.db.Save(&status)
	return s.db.Error
}
