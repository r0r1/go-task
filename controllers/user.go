package controllers

import (
	"encoding/json"
	"go-task/models"
	"strconv"

	"github.com/astaxie/beego"
)

// User
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (this *UserController) Post() {
	var user models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	id, err := models.AddUser(&user)
	if err == nil {
		this.Data["json"] = map[string]int64{"id": id}
		this.Ctx.Output.SetStatus(201)
	} else {
		this.Data["json"] = err.Error()
		this.Ctx.Output.SetStatus(422)
	}
	this.ServeJSON()
}

// @Title Get Detail of User
// @Description get User by id
// @Param id    path  string  true    "The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (this *UserController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	u, err := models.ShowUser(id)
	if err != nil {
		this.Data["json"] = map[string]string{"error": "User Not Found."}
		this.Ctx.Output.SetStatus(404)
	} else {
		this.Data["json"] = u
	}
	this.ServeJSON()
}

// @Title Get All User
// @Description Get All User
// @Param fields  query string  false "Fields returned. e.g. col1,col2 ..."
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (this *UserController) GetAll() {
	var fields []string
	l, err := models.GetAllUser(fields)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param id    path  int  true    "The id you want to update"
// @Param body    body  models.User true    "body for Task content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (this *UserController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUser(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

// @Title Delete User
// @Description delete user
// @Param id    path  string  true    "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *UserController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DestroyUser(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
