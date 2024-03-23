// commentService.go

package services

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserByID(userID uint) (*models.User, error) {
	db := database.GetDB()
	var user models.User

	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func getPhotoByID(photoID uint) (*models.Photo, error) {
	db := database.GetDB()
	var photo models.Photo

	if err := db.Where("id = ?", photoID).First(&photo).Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func GetAllComments() ([]gin.H, error) {
    db := database.GetDB()
    var comments []models.Comment

    if err := db.Find(&comments).Error; err != nil {
        return nil, err
    }

    var response []gin.H
    for _, comment := range comments {
        user, err := getUserByID(comment.UserID)
        if err != nil {
            return nil, err
        }

        photo, err := getPhotoByID(comment.PhotoID)
        if err != nil {
            return nil, err
        }

        commentData := gin.H{
            "id":         comment.ID,
            "message":    comment.Message,
            "photo_id":   comment.PhotoID,
            "user_id":    comment.UserID,
            "created_at": comment.CreatedAt,
            "updated_at": comment.UpdatedAt,
            "User": gin.H{
                "id":       user.ID,
                "email":    user.Email,
                "username": user.Username,
            },
            "photo": gin.H{
                "id":        photo.ID,
                "title":     photo.Title,
                "caption":   photo.Caption,
                "photo_url": photo.PhotoURL,
                "user_id":   photo.UserID,
            },
        }

        response = append(response, commentData)
    }

    return response, nil
}

func CreateComment(c *gin.Context, userID uint) (*models.Comment, error) {
    db := database.GetDB()
    var comment models.Comment
	contentType := helpers.GetContentType(c)

    if contentType == appJSON {
		c.BindJSON(&comment)
	} else {
		c.Bind(&comment)
	}

    comment.UserID = userID

    if err := db.Create(&comment).Error; err != nil {
        return nil, err
    }

    return &comment, nil
}

func UpdateComment(c *gin.Context, userID uint) (*models.Comment, error) {
    db := database.GetDB()
    var comment models.Comment
	contentType := helpers.GetContentType(c)

    commentID, _ := strconv.Atoi(c.Param("commentId"))
    if err := db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
        return nil, err
    }

    if contentType == appJSON {
		c.BindJSON(&comment)
	} else {
		c.Bind(&comment)
	}

    comment.ID = uint(commentID)
    comment.UserID = userID

    if err := db.Save(&comment).Error; err != nil {
        return nil, err
    }

    return &comment, nil
}

func DeleteComment(c *gin.Context, userID uint) error {
    db := database.GetDB()
    commentID, _ := strconv.Atoi(c.Param("commentId"))

    if err := db.Where("id = ? AND user_id = ?", commentID, userID).Delete(&models.Comment{}).Error; err != nil {
        return err
    }

    return nil
}
