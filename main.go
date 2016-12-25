package main

import (
	"go-task/models"
	"go-task/resources"

	"github.com/kataras/iris"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// routes
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello World!")
	})

	// Status
	iris.API("/statuses", resources.UserAPI{})

	iris.Listen(":8080")
}
