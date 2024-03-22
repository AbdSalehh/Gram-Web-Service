package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoGetAll(c *gin.Context) {

	db := database.GetDB()
	Photos := []models.Photo{}

	err := db.Debug().Model(&models.Photo{}).Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}

func PhotoGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photos := []models.Photo{}
	userID := uint(userData["id"].(float64))

	err := db.Debug().Model(&models.Photo{}).Where("user_id = ?", userID).Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, Photos)
}

func PhotoCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.BindJSON(&Photo)
	} else {
		c.Bind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error":   "BAD REQUEST",
				"message": err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func PhotoUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.BindJSON(&Photo)
	} else {
		c.Bind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoURL: Photo.PhotoURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	updatedPhoto := models.Photo{}
	db.First(&updatedPhoto, photoId)

	c.JSON(http.StatusOK, updatedPhoto)
}

func PhotoDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ? AND user_id = ?", photoId, userID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
