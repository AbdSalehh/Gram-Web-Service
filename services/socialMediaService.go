package services

import (
	"MyGram/database"
	"MyGram/models"
)

func GetAllSocialMedia() ([]map[string]interface{}, error) {
	db := database.GetDB()
	var socialMedias []models.SocialMedia

	if err := db.Debug().Find(&socialMedias).Error; err != nil {
		return nil, err
	}

	var socialMediaResponses []map[string]interface{}
	for _, socialMedia := range socialMedias {
		user, err := getUserByID(socialMedia.UserID)
		if err != nil {
			return nil, err
		}

		socialMediaResponse := map[string]interface{}{
			"id":               socialMedia.ID,
			"name":             socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaURL,
			"UserId":           socialMedia.UserID,
			"createdAt":        socialMedia.CreatedAt,
			"updatedAt":        socialMedia.UpdatedAt,
			"User": map[string]interface{}{
				"id":       user.ID,
				"email":    user.Email,
				"username": user.Username,
			},
		}
		socialMediaResponses = append(socialMediaResponses, socialMediaResponse)
	}

	return socialMediaResponses, nil
}

func CreateSocialMedia(socialMedia *models.SocialMedia, userID uint) error {
	db := database.GetDB()
	socialMedia.UserID = userID

	if err := db.Create(socialMedia).Error; err != nil {
		return err
	}

	return nil
}

func UpdateSocialMedia(socialMediaID string, socialMedia *models.SocialMedia, userID uint) error {
	db := database.GetDB()

	socialMedia.UserID = userID

	if err := db.Model(socialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
	}).Error; err != nil {
		return err
	}

	return nil
}

func DeleteSocialMedia(socialMediaID string) error {
	db := database.GetDB()

	if err := db.Where("id = ?", socialMediaID).Delete(&models.SocialMedia{}).Error; err != nil {
		return err
	}

	return nil
}
