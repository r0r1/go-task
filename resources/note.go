package resources

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rorikurniadi/go-task/models"
	"github.com/rorikurniadi/gor-pagination"
)

// NewNoteStorage initializes the storage
func NoteDB(db *gorm.DB) *NoteResource {
	return &NoteResource{db}
}

// Note Resource
type NoteResource struct {
	db *gorm.DB
}

func (nr *NoteResource) Get(c *gin.Context) {
	var notes []models.Note
	var count int
	var noteData []interface{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}

	var pagination = new(libs.Paginate)
	pagination.Page = page
	pagination.PerPage = 15

	nr.db.Preload("User").Order("created_at desc").
		Find(&notes).
		Count(&count)

	pagination.SetTotal(count)

	nr.db.Preload("User").Limit(pagination.PerPage).
		Offset(pagination.Offset()).
		Find(&notes)

	for _, note := range notes {
		json := map[string]interface{}{
			"id":         note.ID,
			"name":       note.Name,
			"user":       note.User.Name,
			"created_at": note.CreatedAt,
		}
		noteData = append(noteData, json)
	}

	c.JSON(200, gin.H{
		"total":    pagination.Total,
		"per_page": pagination.PerPage,
		"prev":     pagination.PrevPage(),
		"next":     pagination.NextPage(),
		"last":     pagination.LastPage(),
		"limit":    pagination.Limit,
		"items":    noteData,
	})
}

func (nr *NoteResource) GetByTask(c *gin.Context) {
	var notes []models.Note
	var count int
	var noteData []interface{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}

	var pagination = new(libs.Paginate)
	pagination.Page = page
	pagination.PerPage = 15

	nr.db.Preload("User").Order("created_at desc").
		Where("task_id = ?", c.Params.ByName("id")).
		Find(&notes).
		Count(&count)

	pagination.SetTotal(count)

	nr.db.Preload("User").Limit(pagination.PerPage).
		Where("task_id = ?", c.Params.ByName("id")).
		Offset(pagination.Offset()).
		Find(&notes)

	for _, note := range notes {
		json := map[string]interface{}{
			"id":         note.ID,
			"name":       note.Name,
			"user":       note.User.Name,
			"created_at": note.CreatedAt,
		}
		noteData = append(noteData, json)
	}

	c.JSON(200, gin.H{
		"total":    pagination.Total,
		"per_page": pagination.PerPage,
		"prev":     pagination.PrevPage(),
		"next":     pagination.NextPage(),
		"last":     pagination.LastPage(),
		"limit":    pagination.Limit,
		"items":    noteData,
	})
}

func (nr *NoteResource) Show(c *gin.Context) {
	id, err := nr.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	var note models.Note

	if nr.db.First(&note, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, note)
	}
}

func (nr *NoteResource) Store(c *gin.Context) {
	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(422, gin.H{"errors": err})
	} else {
		nr.db.Create(&note)
		c.JSON(201, gin.H{"id": note.ID})
	}
}

func (nr *NoteResource) Update(c *gin.Context) {
	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	data := note

	if nr.db.First(&note, c.Params.ByName("id")).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	note.Name = data.Name
	nr.db.Save(&note)

	c.JSON(200, note)
}

func (nr *NoteResource) Destroy(c *gin.Context) {
	var note models.Note

	id, err := nr.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	if nr.db.First(&note, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Tag not found"})
		return
	}

	nr.db.Delete(&note)
	c.JSON(200, gin.H{"message": "Delete note has been successful."})
}

func (nr *NoteResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
