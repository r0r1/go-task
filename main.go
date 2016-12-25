package main

import (
	"github.com/rorikurniadi/go-task/models"
	"github.com/rorikurniadi/go-task/resources"

	"github.com/gin-gonic/gin"
)

// Cors ...
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, PUT, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
		c.Next()
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	statusResource := resources.NewStatusStorage(db)
	authResource := resources.AuthDB(db)
	taskResource := resources.TaskDB(db)

	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/register", authResource.Register)
		v1.POST("/login", authResource.Login().LoginHandler)
		v1.GET("/logout", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Logout",
			})
		})

	}

	auth := r.Group("api/v1")
	auth.Use(authResource.Login().MiddlewareFunc())
	{
		auth.GET("/refresh_token", authResource.Login().RefreshHandler)

		// task
		auth.GET("/tasks", taskResource.Get)
		auth.GET("/tasks/:id", taskResource.Show)
		auth.POST("/tasks", taskResource.Store)
		auth.PUT("/tasks/:id", taskResource.Update)
		auth.DELETE("/tasks/:id", taskResource.Destroy)

		// statuses
		auth.GET("/statuses", statusResource.Get)
		auth.GET("/statuses/:id", statusResource.Show)
		auth.POST("/statuses", statusResource.Store)
	}

	r.Run(":8080")
}
