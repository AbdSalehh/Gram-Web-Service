package controllers

import (
	"MyGram/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoGetAll(c *gin.Context) {
    photos, err := services.GetAllPhotos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Oops.. Something went wrong",
		})
		return
    }

    c.JSON(http.StatusOK, photos)
}

func PhotoCreate(c *gin.Context) {
    userID := getUserIDFromContext(c)
    photo, err := services.CreatePhoto(c, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

	c.JSON(http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoURL,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	})
}

func PhotoUpdate(c *gin.Context) {
    userID := getUserIDFromContext(c)
    photo, err := services.UpdatePhoto(c, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoURL,
		"user_id":    photo.UserID,
		"updated_at": photo.UpdatedAt,
	})
}

func PhotoDelete(c *gin.Context) {
    userID := getUserIDFromContext(c)
    if err := services.DeletePhoto(c, userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Oops.. Something went wrong",
		})
		return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Your photo has been successfully deleted",
    })
}

func getUserIDFromContext(c *gin.Context) uint {
    userData := c.MustGet("userData").(jwt.MapClaims)
    return uint(userData["id"].(float64))
}
