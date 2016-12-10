package main

import (
    "fmt"
    _ "go-task/routers"
    "time"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/plugins/cors"
    _ "github.com/go-sql-driver/mysql"
    "go-task/models"
)

func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)

    dbUser := beego.AppConfig.String("db_user")
    dbPassword := beego.AppConfig.String("db_password")
    dbName := beego.AppConfig.String("db_name")
    dsn := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8"

    orm.RegisterDataBase("default", "mysql", dsn)
    orm.DefaultTimeLoc = time.UTC
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

    // Database alias.
    name := "default"
    // Drop table and re-create.
    force := true
    // Print log.
    verbose := true
    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        fmt.Println(err)
    }

    // Create Status Default
    o := orm.NewOrm()
    statusPending := new(models.TaskStatus)
    statusPending.Name = "Pending"
    statusPending.Label = "label label-info"

    statusDone := new(models.TaskStatus)
    statusDone.Name = "Done"
    statusDone.Label = "label label-success"

    statusProgress := new(models.TaskStatus)
    statusProgress.Name = "Progress"
    statusProgress.Label = "label label-warning"

    o.Insert(statusPending)
    o.Insert(statusDone)
    o.Insert(statusProgress)

    beego.Run()
    orm.RunCommand()
}
