package resources

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rorikurniadi/go-task/libs"
	"github.com/rorikurniadi/go-task/models"
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
	var count int
	var taskData []interface{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}

	var pagination = new(libs.Paginate)
	pagination.Page = page
	pagination.PerPage = 15

	tr.db.Preload("User").Order("created_at desc").
		Find(&tasks).
		Count(&count)

	pagination.SetTotal(count)

	tr.db.Preload("User").Limit(pagination.PerPage).
		Offset(pagination.Offset()).
		Find(&tasks)

	for _, task := range tasks {
		json := map[string]interface{}{
			"id":          task.ID,
			"name":        task.Name,
			"priority":    task.Priority,
			"status":      task.Status,
			"description": task.Description,
			"date":        task.CreatedAt,
			"user":        task.User.Name,
		}
		taskData = append(taskData, json)
	}

	c.JSON(200, gin.H{
		"total":    pagination.Total,
		"per_page": pagination.PerPage,
		"prev":     pagination.PrevPage(),
		"next":     pagination.NextPage(),
		"last":     pagination.LastPage(),
		"limit":    pagination.Limit,
		"items":    taskData,
	})
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
