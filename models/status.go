package models

type Status struct {
    Name  string `valid:"Required"`
    Label string `gorm:"type:text"`
    Task  []Task
}
