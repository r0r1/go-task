package models

import (
    "github.com/astaxie/beego/orm"
)

type TaskStatus struct {
    Id    int
    Name  string `valid:"Required"`
    Label string `orm:"null"`
    Task  *Task  `orm:"reverse(one)"`
}

func (t *TaskStatus) TableName() string {
    return "tasks_status"
}

func init() {
    orm.RegisterModel(new(TaskStatus))
}
