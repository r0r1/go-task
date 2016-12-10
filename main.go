package main

import (
    "fmt"
    _ "go-task/routers"
    "time"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/plugins/cors"
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

    // Handle CORS
    beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"GET", "POST", "PUT", "PATCH"},
        AllowHeaders: []string{"Origin"},
        ExposeHeaders: []string{"Content-Length"},
        AllowCredentials: true,
    }))

    beego.Run()
    orm.RunCommand()
}
