package database

import (
	"MyGram/models"
)

func UpdateUser(user models.User) error {
    db := GetDB()
    return db.Debug().Model(&user).Where("id = ?", user.ID).Updates(models.User{
        Username: user.Username,
        Email:    user.Email,
    }).Error
}

func DeleteUser(userID uint) error {
    db := GetDB()
    tx := db.Begin()

    if err := tx.Where("user_id = ?", userID).Delete(&models.SocialMedia{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Where("photo_id IN (SELECT id FROM photos WHERE user_id = ?)", userID).Delete(&models.Comment{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Where("user_id = ?", userID).Delete(&models.Photo{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    if err := tx.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    return nil
}
