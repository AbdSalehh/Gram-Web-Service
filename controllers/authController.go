package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": "Bad Request",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id": User.ID,
		"username": User.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": "Oops! Something went wrong.",
        })
        return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, uint(User.Age), User.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}