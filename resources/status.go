package resources

import (
	"go-task/models"
	"go-task/services/domain"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

// Status API
type StatusAPI struct {
	ctx *iris.Context
	db  *gorm.DB
}

// Status Storage
type StatusStorage struct {
	db *gorm.DB
}

// GET /statuses
func (s StatusAPI) Get() {
	var statuses []models.Status
	s.db.Order("id").Find(&statuses)
	return statuses
	// statuses := db.Find(&Status{})
	// c.JSON(iris.StatusOK, statuses)
}

// GET /:param1 which its value passed to the id argument
func (s StatusAPI) GetBy(id string) { // id equals to u.Param("param1")
	s.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))
}

// PUT /statuses
func (s StatusAPI) Put() {
	name := s.FormValue("name")
	// myDb.InsertUser(...)
	println(string(name))
	println("Put from /users")
}

// POST /statuses/:param1
func (s StatusAPI) PostBy(id string) {

	name := s.PostValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)

	statusCustom := models.Status{
		Name:  "Custom",
		Label: "label label-custom",
	}

	id := *domain.StatusStorage.Insert(statusCustom)

	s.JSON(iris.StatusOK, iris.Map{"message": "add status successfull."})
	s.SetStatusCode(201)

	println(string(name))
	println("Post from /users/" + id)
}

// DELETE /statuses/:param1
func (s StatusAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}
