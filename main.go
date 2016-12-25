package main

import (
	"go-task/models"
	"go-task/resources"

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
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/register", authResource.Register)
		v1.POST("/login", authResource.Login)
		v1.GET("/logout", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Login",
			})
		})

		v1.GET("/statuses", statusResource.Get)
		v1.GET("/statuses/:id", statusResource.Show)
		v1.POST("/statuses", statusResource.Store)
	}

	r.Run(":8080")
}

func jwtMiddleware() {

}
