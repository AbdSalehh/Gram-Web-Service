package controllers

import (
	"MyGram/helpers"
	"MyGram/models"
	"MyGram/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediaGetAll(c *gin.Context) {
	socialMedias, err := services.GetAllSocialMedia()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"social_medias": socialMedias})
}

func SocialMediaCreate(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	if err := services.CreateSocialMedia(&socialMedia, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	})
}

func SocialMediaUpdate(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	socialMediaID := c.Param("socialMediaId")

	socialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&socialMedia)
	} else {
		c.ShouldBind(&socialMedia)
	}

	if err := services.UpdateSocialMedia(socialMediaID, &socialMedia, userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":               socialMediaID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.UserID,
		"updated_at":       socialMedia.UpdatedAt,
	})
}

func SocialMediaDelete(c *gin.Context) {
	socialMediaID := c.Param("socialMediaId")

	if err := services.DeleteSocialMedia(socialMediaID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
