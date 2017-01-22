package resources

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rorikurniadi/go-task/models"
	"github.com/rorikurniadi/gor-pagination"
)

// NewTagStorage initializes the storage
func TagDB(db *gorm.DB) *TagResource {
	return &TagResource{db}
}

// Tag Resource
type TagResource struct {
	db *gorm.DB
}

func (tg *TagResource) Get(c *gin.Context) {
	var tags []models.Tag
	var count int
	var tagData []interface{}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 0
	}

	var pagination = new(libs.Paginate)
	pagination.Page = page
	pagination.PerPage = 15

	tg.db.Order("created_at desc").
		Find(&tags).
		Count(&count)

	pagination.SetTotal(count)

	tg.db.Limit(pagination.PerPage).
		Offset(pagination.Offset()).
		Find(&tags)

	for _, tag := range tags {
		json := map[string]interface{}{
			"id":   tag.ID,
			"name": tag.Name,
		}
		tagData = append(tagData, json)
	}

	c.JSON(200, gin.H{
		"total":    pagination.Total,
		"per_page": pagination.PerPage,
		"prev":     pagination.PrevPage(),
		"next":     pagination.NextPage(),
		"last":     pagination.LastPage(),
		"limit":    pagination.Limit,
		"items":    tagData,
	})
}

func (tg *TagResource) Show(c *gin.Context) {
	id, err := tg.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	var tag models.Tag

	if tg.db.First(&tag, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, tag)
	}
}

func (tg *TagResource) Store(c *gin.Context) {
	var tag models.Tag

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(422, gin.H{"errors": err})
	} else {
		tg.db.Create(&tag)
		c.JSON(201, gin.H{"id": tag.ID})
	}
}

func (tg *TagResource) Update(c *gin.Context) {
	var tag models.Tag

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(422, gin.H{"errors": err})
		return
	}

	data := tag

	if tg.db.First(&tag, c.Params.ByName("id")).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	tag.Name = data.Name
	tg.db.Save(&tag)

	c.JSON(200, tag)
}

func (tg *TagResource) Destroy(c *gin.Context) {
	var tag models.Tag

	id, err := tg.getId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "problem decoding id sent"})
		return
	}

	if tg.db.First(&tag, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "Tag not found"})
		return
	}

	tg.db.Delete(&tag)
	c.JSON(200, gin.H{"message": "Delete tag has been successful."})
}

func (tr *TagResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
