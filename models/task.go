package models

import (
    "github.com/jinzhu/gorm"
)

type Task struct {
    gorm.Model
    Name        string `valid:"Required"`
    UserId      uint
    Description string `gorm:"type:text"`
    Parent      *Task  `valid:"Required"`
    Priority    int    `valid:"Min:1;Max:5"`
    StatusId    uint
}
