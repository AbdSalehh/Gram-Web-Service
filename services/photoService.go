package services

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPhotos() ([]gin.H, error) {
    db := database.GetDB()
    var photos []models.Photo
    if err := db.Find(&photos).Error; err != nil {
        return nil, err
    }

    var photoResponses []gin.H
    for _, photo := range photos {
        user, err := getUserByPhotoID(photo.UserID)
        if err != nil {
            return nil, err
        }

        photoResponse := gin.H{
            "id":         photo.ID,
            "title":      photo.Title,
            "caption":    photo.Caption,
            "photo_url":  photo.PhotoURL,
            "user_id":    photo.UserID,
            "user": gin.H{
                "email":    user.Email,
                "username": user.Username,
            },
            "created_at": photo.CreatedAt,
            "updated_at": photo.UpdatedAt,
        }
        photoResponses = append(photoResponses, photoResponse)
    }

    return photoResponses, nil
}

func CreatePhoto(c *gin.Context, userID uint) (*models.Photo, error) {
    db := database.GetDB()
    var photo models.Photo
	contentType := helpers.GetContentType(c)

    if contentType == appJSON {
		c.BindJSON(&photo)
	} else {
		c.Bind(&photo)
	}

    photo.UserID = userID

    if err := db.Create(&photo).Error; err != nil {
        return nil, err
    }

    return &photo, nil
}

func UpdatePhoto(c *gin.Context, userID uint) (*models.Photo, error) {
    db := database.GetDB()
    var photo models.Photo
	contentType := helpers.GetContentType(c)

    photoID, _ := strconv.Atoi(c.Param("photoId"))
    if err := db.Where("id = ? AND user_id = ?", photoID, userID).First(&photo).Error; err != nil {
        return nil, err
    }

    if contentType == appJSON {
		c.BindJSON(&photo)
	} else {
		c.Bind(&photo)
	}


    photo.ID = uint(photoID)
    photo.UserID = userID

    if err := db.Save(&photo).Error; err != nil {
        return nil, err
    }

    return &photo, nil
}

func DeletePhoto(c *gin.Context, userID uint) error {
    db := database.GetDB()
    photoID, _ := strconv.Atoi(c.Param("photoId"))

    if err := db.Where("id = ? AND user_id = ?", photoID, userID).Delete(&models.Photo{}).Error; err != nil {
        return err
    }

    return nil
}

func getUserByPhotoID(userID uint) (*models.User, error) {
    db := database.GetDB()
    var user models.User
    if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
