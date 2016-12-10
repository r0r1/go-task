package models

import (
    "fmt"

    "github.com/astaxie/beego/orm"
)

type TaskStatus struct {
    Id    int
    Name  string `valid:"Required"`
    Label string `orm:"null"`
}

func (t *TaskStatus) TableName() string {
    return "tasks_status"
}

func init() {
    orm.RegisterModel(new(TaskStatus))
}

func AddStatus(status *TaskStatus) (id int64, err error) {
    o := orm.NewOrm()
    id, err = o.Insert(status)
    if err != nil {
        return 0, err
    }

    return id, nil
}

func ShowStatus(id int) (status *TaskStatus, err error) {
    o := orm.NewOrm()
    status = &TaskStatus{Id: id}
    if err = o.Read(status); err == nil {
        return status, nil
    }
    return nil, err
}

func GetAllStatus(fields []string) (ml []interface{}, err error) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(TaskStatus))
    var l []TaskStatus

    if _, err := qs.All(&l); err == nil {
        for _, v := range l {
            ml = append(ml, v)
        }
        return ml, nil
    }
    return nil, err
}

func UpdateStatus(ts *TaskStatus) (err error) {
    o := orm.NewOrm()
    v := TaskStatus{Id: ts.Id}
    // ascertain id exists in the database
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Update(ts); err == nil {
            fmt.Println("Number of records updated in database:", num)
        }
    }
    return
}

func DestroyStatus(id int) (err error) {
    o := orm.NewOrm()
    v := TaskStatus{Id: id}
    // ascertain id exists in the database
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Delete(&TaskStatus{Id: id}); err == nil {
            fmt.Println("Number of records deleted in database:", num)
        }
    }
    return
}
