package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentGetAll(c *gin.Context) {

	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Debug().Model(&models.Comment{}).Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, Comments)
}

func CommentGet(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comments := []models.Comment{}
	userID := uint(userData["id"].(float64))

	//return all comments by user
	err := db.Debug().Model(&models.Comment{}).Where("user_id = ?", userID).Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, Comments)

}

func CommentCreate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	errPhoto := db.First(&Photo, Comment.PhotoID).Error
	if errPhoto != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": "photo not found",
		})
		return
	}

	Comment.UserID = userID

	errComment := db.Create(&Comment).Error
	if errComment != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": errComment.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)

}

func CommentUpdate(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	commentID := c.Param("commentId")

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", userID).Updates(
		models.Comment{
			Message: Comment.Message,
		},
	).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	updatedComment := models.Comment{}
	db.First(&updatedComment, commentID)

	c.JSON(http.StatusOK, gin.H{
		"id":         updatedComment.ID,
		"message":    updatedComment.Message,
		"photo_id":   updatedComment.PhotoID,
		"updated_at": updatedComment.UpdatedAt,
		"user_id":    updatedComment.UserID,
	})

}

func CommentDelete(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	commentID := c.Param("commentId")
	userID := uint(userData["id"].(float64))

	err := db.Where("id = ? AND user_id = ?", commentID, userID).Delete(&models.Comment{}).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "NOT FOUND",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,
		gin.H{
			"message": "Your comment has been successfully deleted",
		},
	)

}
