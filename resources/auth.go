package resources

import (
	"go-task/models"

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

func (ar *AuthResource) Login(c *gin.Context) {
	var login Login
	var user models.User

	if err := c.BindJSON(&login); err != nil {
		c.JSON(422, gin.H{"error": "email and password are required"})
		return
	}

	data := ar.db.Where("email = ?", login.Email).Find(&user)

	if data.RecordNotFound() {
		c.JSON(401, gin.H{"error": "Login failed, email not found."})
		return
	}

	hashedPassword := []byte(user.Password)
	password := []byte(login.Password)

	checkPassword := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if checkPassword != nil {
		c.JSON(401, gin.H{"error": "Login failed, wrong password."})
		return
	}

	c.JSON(200, data.Value)
}
