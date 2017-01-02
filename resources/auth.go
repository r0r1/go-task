package resources

import (
	"time"

	"github.com/rorikurniadi/go-task/models"

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
	// JWT Middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "go-task-management",
		Key:        []byte("go-task-management"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			var user = new(models.User)
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

func (ar *AuthResource) CurrentUser(c *gin.Context) {
	var user = new(models.User)
	jwtClaims := jwt.ExtractClaims(c)
	id := jwtClaims["id"]
	data := ar.db.Where("email = ?", id).Find(&user)

	if data.RecordNotFound() {
		c.JSON(404, gin.H{
			"code":    404,
			"message": "User not found.",
		})
	}

	c.JSON(200, user)
}

func (ar *AuthResource) Get(c *gin.Context) {
	var users []models.User

	ar.db.Order("created_at desc").Find(&users)

	c.JSON(200, users)
}
