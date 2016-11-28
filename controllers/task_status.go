package controllers

import (
    "encoding/json"
    "go-task/models"
    "strconv"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
)

// Status Task
type TaskStatusController struct {
    beego.Controller
}

func (this *TaskStatusController) URLMapping() {
    this.Mapping("Post", this.Post)
    this.Mapping("GetOne", this.GetOne)
    this.Mapping("GetAll", this.GetAll)
    this.Mapping("Put", this.Put)
    this.Mapping("Delete", this.Delete)
}

// @Title Create Task Status
// @Description Create Task Task
// @Param body    body  models.TaskStatus true    "body for Task  Status content"
// @Success 200 {int} models.TaskStatus.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TaskStatusController) Post() {
    var m models.TaskStatus
    json.Unmarshal(this.Ctx.Input.RequestBody, &m)

    valid := validation.Validation{}
    valid.Valid(&m)
    if valid.HasErrors() {
        errorMessages := make(map[string]string)
        for _, err := range valid.Errors {
            errorMessages[err.Key] = err.Message
        }
        this.Data["json"] = errorMessages
        this.Ctx.Output.SetStatus(422)
        this.ServeJSON()
        return
    }

    id, err := models.AddStatus(&m)
    if err == nil {
        this.Data["json"] = map[string]int64{"id": id}
        this.Ctx.Output.SetStatus(201)
    } else {
        this.Data["json"] = err.Error()
        this.Ctx.Output.SetStatus(422)
    }
    this.ServeJSON()
}

// @Title Get Detail Status
// @Description get Task Status by id
// @Param id    path  string  true    "The key for staticblock"
// @Success 200 {object} models.TaskStatus
// @Failure 403 :id is empty
// @router /:id [get]
func (this *TaskStatusController) GetOne() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    t, err := models.ShowStatus(id)
    if err != nil {
        this.Data["json"] = map[string]string{"error": "Status Not Found."}
        this.Ctx.Output.SetStatus(404)
    } else {
        this.Data["json"] = t
    }
    this.ServeJSON()
}

// @Title Get All
// @Description get Task Status
// @Param fields  query string  false "Fields returned. e.g. col1,col2 ..."
// @Success 200 {object} models.TaskStatus
// @Failure 403
// @router / [get]
func (this *TaskStatusController) GetAll() {
    var fields []string
    l, err := models.GetAllStatus(fields)
    if err != nil {
        this.Data["json"] = err.Error()
    } else {
        this.Data["json"] = l
    }
    this.ServeJSON()
}

// @Title Update
// @Description update the Task Status
// @Param id    path  string  true    "The id you want to update"
// @Param body    body  models.TaskStatus true    "body for Task content"
// @Success 200 {object} models.TaskStatus
// @Failure 403 :id is not int
// @router /:id [put]
func (this *TaskStatusController) Put() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    v := models.TaskStatus{Id: id}
    json.Unmarshal(this.Ctx.Input.RequestBody, &v)
    if err := models.UpdateStatus(&v); err == nil {
        this.Data["json"] = "OK"
    } else {
        this.Data["json"] = err.Error()
    }
    this.ServeJSON()
}

// @Title Delete
// @Description delete the Task
// @Param id    path  string  true    "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *TaskStatusController) Delete() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    if err := models.DestroyStatus(id); err == nil {
        this.Data["json"] = "OK"
    } else {
        this.Data["json"] = err.Error()
    }
    this.ServeJSON()
}
