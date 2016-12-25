package resources

import (
	"log"
	"strconv"

	"github.com/rorikurniadi/go-task/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NewUserStorage initializes the storage
func TaskDB(db *gorm.DB) *TaskResource {
	return &TaskResource{db}
}

// Status API
type TaskResource struct {
	db *gorm.DB
}

func (tr *TaskResource) Get(c *gin.Context) {
	var tasks []models.Task

	tr.db.Order("created_at desc").Find(&tasks)

	c.JSON(200, tasks)
}

func (tr *TaskResource) Show(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	var task models.Task

	if tr.db.First(&task, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, task)
	}
}

func (tr *TaskResource) Store(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(422, gin.H{"errors": err})
	} else {
		tr.db.Create(&task)
		c.JSON(201, gin.H{"id": task.ID})
	}
}

func (tr *TaskResource) Update(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	data := task

	if tr.db.First(&task, c.Params.ByName("id")).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	task.Name = data.Name
	task.Parent = data.Parent
	task.Priority = data.Priority
	task.Status = data.Status
	task.Description = data.Description

	tr.db.Save(&task)

	c.JSON(200, task)
}

func (tr *TaskResource) Destroy(c *gin.Context) {
	var task models.Task

	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	if tr.db.First(&task, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	tr.db.Delete(&task)
	c.JSON(200, gin.H{"message": "Delete task has been successful."})
}

func (tr *TaskResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
