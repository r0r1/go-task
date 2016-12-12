package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

type User struct {
	Id       int
	Name     string    `valid:"Required"`
	Email    string    `orm:"unique" valid:"Required;Email"`
	Username string    `orm:"unique" valid:"Required"`
	Password string    `valid:"Required"`
	Contact  string    `orm:"null"`
	Address  string    `orm:"null;type(text)"`
	Task 	[]*Task    `orm:"reverse(many)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

func (t *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u *User) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(u)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func ShowUser(id int) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Id: id}
	if err = o.Read(user); err == nil {
		return user, nil
	}
	return nil, err
}

func GetAllUser(fields []string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	var l []User

	if _, err := qs.All(&l); err == nil {
		for _, v := range l {
			ml = append(ml, v)
		}
		return ml, nil
	}
	return nil, err
}

func UpdateUser(u *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: u.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(u); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DestroyUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
