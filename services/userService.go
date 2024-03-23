package services

import (
	"MyGram/database"
	"MyGram/models"
)

var (
	appJSON = "application/json"
)

func UpdateUserService(user models.User) error {
    return database.UpdateUser(user)
}

func DeleteUserService(userID uint) error {
    return database.DeleteUser(userID)
}
