package controllers

import (
    "encoding/json"
    "errors"
    "go-task/models"
    "strconv"
    "strings"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/validation"
)

// Task
type TaskController struct {
    beego.Controller
}

func (this *TaskController) URLMapping() {
    this.Mapping("Post", this.Post)
    this.Mapping("GetOne", this.GetOne)
    this.Mapping("GetAll", this.GetAll)
    this.Mapping("Put", this.Put)
    this.Mapping("Delete", this.Delete)
}

// @Title Create Task
// @Description Create Task
// @Param body    body  models.Task true    "body for Task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (this *TaskController) Post() {
    var v models.Task
    json.Unmarshal(this.Ctx.Input.RequestBody, &v)

    valid := validation.Validation{}
    valid.Valid(&v)
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
    beego.Info(&v)
    id, err := models.StoreTask(&v)
    if err == nil {
        this.Data["json"] = map[string]int64{"id": id}
        this.Ctx.Output.SetStatus(201)
    } else {
        this.Data["json"] = err.Error()
        this.Ctx.Output.SetStatus(422)
    }
    this.ServeJSON()
}

// @Title Get
// @Description get Task by id
// @Param id    path  string  true    "The key for staticblock"
// @Success 200 {object} models.Task
// @Failure 403 :id is empty
// @router /:id [get]
func (this *TaskController) GetOne() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    t, err := models.ShowTask(id)
    if err != nil {
        this.Data["json"] = map[string]string{"error": "Task Not Found."}
        this.Ctx.Output.SetStatus(404)
    } else {
        this.Data["json"] = t
    }
    this.ServeJSON()
}

// @Title Get All
// @Description get Task
// @Param query query string  false "Filter. e.g. col1:v1,col2:v2 ..."
// @Param fields  query string  false "Fields returned. e.g. col1,col2 ..."
// @Param sortby  query string  false "Sorted-by fields. e.g. col1,col2 ..."
// @Param order query string  false "Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param limit query string  false "Limit the size of result set. Must be an integer"
// @Param offset  query string  false "Start position of result set. Must be an integer"
// @Success 200 {object} models.Task
// @Failure 403
// @router / [get]
func (this *TaskController) GetAll() {
    var fields []string
    var sortby []string
    var order []string
    var query map[string]string = make(map[string]string)
    var limit int64 = 10
    var offset int64 = 0

    // fields: col1,col2,entity.col3
    if v := this.GetString("fields"); v != "" {
        fields = strings.Split(v, ",")
    }
    // limit: 10 (default is 10)
    // if v, err := this.GetInt("limit"); err == nil {
    //  limit = v
    // }
    // // offset: 0 (default is 0)
    // if v, err := this.GetInt("offset"); err == nil {
    //  offset = v
    // }
    // sortby: col1,col2
    if v := this.GetString("sortby"); v != "" {
        sortby = strings.Split(v, ",")
    }
    // order: desc,asc
    if v := this.GetString("order"); v != "" {
        order = strings.Split(v, ",")
    }
    // query: k:v,k:v
    if v := this.GetString("query"); v != "" {
        for _, cond := range strings.Split(v, ",") {
            kv := strings.Split(cond, ":")
            if len(kv) != 2 {
                this.Data["json"] = errors.New("Error: invalid query key/value pair")
                this.ServeJSON()
                return
            }
            k, v := kv[0], kv[1]
            query[k] = v
        }
    }

    l, err := models.GetAllTask(query, fields, sortby, order, offset, limit)
    if err != nil {
        this.Data["json"] = err.Error()
    } else {
        this.Data["json"] = l
    }
    this.ServeJSON()
}

// @Title Update
// @Description update the Task
// @Param id    path  string  true    "The id you want to update"
// @Param body    body  models.Task true    "body for Task content"
// @Success 200 {object} models.Task
// @Failure 403 :id is not int
// @router /:id [put]
func (this *TaskController) Put() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    v := models.Task{Id: id}
    json.Unmarshal(this.Ctx.Input.RequestBody, &v)
    if err := models.UpdateTask(&v); err == nil {
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
func (this *TaskController) Delete() {
    idStr := this.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    if err := models.DestroyTask(id); err == nil {
        this.Data["json"] = "OK"
    } else {
        this.Data["json"] = err.Error()
    }
    this.ServeJSON()
}
