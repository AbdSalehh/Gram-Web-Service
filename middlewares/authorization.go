package middlewares

import (
	"MyGram/database"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)

		userID := uint(userData["id"].(float64))
		User := models.User{}

		err := db.Select("id").First(&User, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Uppss.. User not found",
			})
			return
		}

		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		photoID, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "INVALID Request",
				"message": "Parameter is not valid",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		photo := models.Photo{}

		err = db.Select("user_id").First(&photo, uint(photoID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Uppss.. Photo not found",
			})
			return
		}

		if photo.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Access Denied",
				"message": "You are not authorized to perform this action",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		commentID, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "INVALID Request",
				"message": "Parameter is not valid",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Uppss.. Comment not found",
			})
			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Access Denied",
				"message": "You are not authorized to perform this action",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "INVALID Request",
				"message": "Parameter is not valid",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		err = db.Select("user_id").First(&socialMedia, uint(socialMediaID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Not Found",
				"message": "Uppss.. Social media not found",
			})
			return
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":   "Access Denied",
				"message": "You are not authorized to perform this action",
			})
			return
		}

		c.Next()
	}
}
