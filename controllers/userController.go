package controllers

import (
	"MyGram/helpers"
	"MyGram/models"
	"MyGram/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
    userData := c.MustGet("userData").(jwt.MapClaims)
    contentType := helpers.GetContentType(c)
    userID := uint(userData["id"].(float64))
    userAge := uint8(userData["age"].(float64))

    var user models.User
    if contentType == appJSON {
        c.ShouldBindJSON(&user)
    } else {
        c.ShouldBind(&user)
    }

    user.ID = userID
    user.Age = userAge

    if err := services.UpdateUserService(user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

	c.JSON(http.StatusCreated, gin.H{
		"age":      user.Age,
		"email":    user.Email,
		"id": user.ID,
		"username": user.Username,
	})
}

func DeleteUser(c *gin.Context) {
    userData := c.MustGet("userData").(jwt.MapClaims)
    userID := uint(userData["id"].(float64))

    if err := services.DeleteUserService(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Your account has been successfully deleted",
    })
}
