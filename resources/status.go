package resources

import (
	"log"
	"strconv"

	"github.com/rorikurniadi/go-task/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserStorage initializes the storage
func NewStatusStorage(db *gorm.DB) *StatusResource {
	return &StatusResource{db}
}

// Status API
type StatusResource struct {
	db *gorm.DB
}

func (sr *StatusResource) Get(c *gin.Context) {
	var statuses []models.Status

	sr.db.Order("created_at desc").Find(&statuses)

	c.JSON(200, statuses)
}

func (sr *StatusResource) Show(c *gin.Context) {
	id, err := sr.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	var status models.Status

	if sr.db.First(&status, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, status)
	}
}

func (sr *StatusResource) Store(c *gin.Context) {
	var status models.Status

	if err := c.BindJSON(&status); err != nil {
		c.JSON(422, gin.H{"errors": err})
	} else {
		sr.db.Create(&status)
		c.JSON(201, gin.H{"id": status.ID})
	}
}

func (sr *StatusResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
