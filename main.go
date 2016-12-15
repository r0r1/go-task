package main

import (
    "fmt"
    _ "go-task/routers"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/plugins/cors"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "go-task/models"
)

func init() {
    dbUser := beego.AppConfig.String("db_user")
    dbPassword := beego.AppConfig.String("db_password")
    dbName := beego.AppConfig.String("db_name")
    dsn := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8"

    db, err := gorm.Open("mysql", dsn)

    if err != nil {
        fmt.Println("error connect")
    }
}

func main() {
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
        beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }

    // Handle CORS
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    db.AutoMigrate(&models.User{}, &models.Task{}, &models.Status{})

    // Create Status Default
    seedData()

    beego.Run()
}

func seedData() {
    statusPending := &models.Status{}
    statusPending.Name = "Pending"
    statusPending.Label = "label label-info"
    db.Create(statusPending)

    statusDone := models.Status
    statusDone.Name = "Done"
    statusDone.Label = "label label-success"
    db.Create(&statusDone)

    statusProgress := models.Status
    statusProgress.Name = "Progress"
    statusProgress.Label = "label label-warning"
    db.Create(&statusProgress)

    user := models.User
    user.Name = "John Doe"
    user.Email = "john@doe.com"
    user.Username = "john-doe"
    user.Password = "foobarbaz"
    db.Create(&user)

    task := models.Task
    task.Name = "Taks 1"
    task.User = user1
    task.Status = statusDone
    task.Priority = 3
    task.Description = "Lorem Ipsum....."
    db.Create(&task)

    taskChild := models.Task
    taskChild.Name = "Taks Child 1"
    taskChild.Parent = task
    taskChild.User = user
    taskChild.Status = statusPending
    taskChild.Priority = 4
    taskChild.Description = "Lorem Ipsum....."
    db.Create(&taskChild)
}
