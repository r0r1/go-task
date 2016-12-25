package resources

import (
	"go-task/models"
	"time"

	jwt "gopkg.in/appleboy/gin-jwt.v2"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Auth DB initializes the storage
func AuthDB(db *gorm.DB) *AuthResource {
	return &AuthResource{db}
}

// Auth Resource
type AuthResource struct {
	db *gorm.DB
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ar *AuthResource) Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(422, gin.H{"errors": err})
	} else {
		password := []byte(user.Password)

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		user.Password = string(hashedPassword)
		ar.db.Save(&user)
		c.JSON(201, user)
	}
}

func (ar *AuthResource) Login() *jwt.GinJWTMiddleware {
	var user models.User

	// JWT Middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "go-task-management",
		Key:        []byte(user.Email + user.Password),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			data := ar.db.Where("email = ?", email).Find(&user)

			if data.RecordNotFound() {
				return email, false
			}

			hashedPassword := []byte(user.Password)
			inputPassword := []byte(password)

			checkPassword := bcrypt.CompareHashAndPassword(hashedPassword, inputPassword)
			if checkPassword != nil {
				return email, false
			}

			return email, true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
	}

	return authMiddleware
}

func (ar *AuthResource) Logout(c *gin.Context) {
	// auth := &jwt.GinJWTMiddleware{
	// 	Timeout: time.Now,
	// 	Unauthorized: func(c *gin.Context, code int, message string) {
	// 		c.JSON(code, gin.H{
	// 			"code":    code,
	// 			"message": message,
	// 		})
	// 	},
	// }
	// return auth
}
