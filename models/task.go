package models

import (
    "errors"
    "fmt"
    "reflect"
    "strings"
    "time"

    "github.com/astaxie/beego/orm"
)

type Task struct {
    Id          int
    Name        string `valid:"Required"`
    Description string `orm:"null;type(text)" valid:"Required"`
    Parent      int    `orm:"null"`
    Priority    int
    Created     time.Time `orm:"auto_now_add;type(datetime)"`
}

func (t *Task) TableName() string {
    return "tasks"
}

func init() {
    orm.RegisterModel(new(Task))
}

func StoreTask(t *Task) (id int64, err error) {
    o := orm.NewOrm()

    id, err = o.Insert(t)
    if err != nil {
        return 0, errors.New("Failed insert data")
    }

    return id, nil
}

// GetTaskById retrieves Task by Id. Returns error if
// Id doesn't exist
func ShowTask(id int) (v *Task, err error) {
    o := orm.NewOrm()
    v = &Task{Id: id}
    if err = o.Read(v); err == nil {
        return v, nil
    }
    return nil, err
}

// GetAllTask retrieves all Task matches certain condition. Returns empty list if
// no records exist
func GetAllTask(query map[string]string, fields []string, sortby []string, order []string,
    offset int64, limit int64) (ml []interface{}, err error) {
    o := orm.NewOrm()
    qs := o.QueryTable(new(Task))
    // query k=v
    for k, v := range query {
        // rewrite dot-notation to Task__Attribute
        k = strings.Replace(k, ".", "__", -1)
        qs = qs.Filter(k, v)
    }
    // order by:
    var sortFields []string
    if len(sortby) != 0 {
        if len(sortby) == len(order) {
            // 1) for each sort field, there is an associated order
            for i, v := range sortby {
                orderby := ""
                if order[i] == "desc" {
                    orderby = "-" + v
                } else if order[i] == "asc" {
                    orderby = v
                } else {
                    return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
                }
                sortFields = append(sortFields, orderby)
            }
            qs = qs.OrderBy(sortFields...)
        } else if len(sortby) != len(order) && len(order) == 1 {
            // 2) there is exactly one order, all the sorted fields will be sorted by this order
            for _, v := range sortby {
                orderby := ""
                if order[0] == "desc" {
                    orderby = "-" + v
                } else if order[0] == "asc" {
                    orderby = v
                } else {
                    return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
                }
                sortFields = append(sortFields, orderby)
            }
        } else if len(sortby) != len(order) && len(order) != 1 {
            return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
        }
    } else {
        if len(order) != 0 {
            return nil, errors.New("Error: unused 'order' fields")
        }
    }

    var l []Task
    qs = qs.OrderBy(sortFields...)
    if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
        if len(fields) == 0 {
            for _, v := range l {
                ml = append(ml, v)
            }
        } else {
            // trim unused fields
            for _, v := range l {
                m := make(map[string]interface{})
                val := reflect.ValueOf(v)
                for _, fname := range fields {
                    m[fname] = val.FieldByName(fname).Interface()
                }
                ml = append(ml, m)
            }
        }
        return ml, nil
    }
    return nil, err
}

// UpdateTask updates Task by Id and returns error if
// the record to be updated doesn't exist
func UpdateTask(t *Task) (err error) {
    o := orm.NewOrm()
    v := Task{Id: t.Id}
    // ascertain id exists in the database
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Update(t); err == nil {
            fmt.Println("Number of records updated in database:", num)
        }
    }
    return
}

// DeleteTask deletes Task by Id and returns error if
// the record to be deleted doesn't exist
func DestroyTask(id int) (err error) {
    o := orm.NewOrm()
    v := Task{Id: id}
    // ascertain id exists in the database
    if err = o.Read(&v); err == nil {
        var num int64
        if num, err = o.Delete(&Task{Id: id}); err == nil {
            fmt.Println("Number of records deleted in database:", num)
        }
    }
    return
}
