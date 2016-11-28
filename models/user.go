package models

import (
	"errors"
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
		return 0, errors.New("Failed insert data")
	}

	return id, nil
}
