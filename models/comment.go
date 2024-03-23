package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func ValidateMessage(message string) error {
    if message == "" {
        return errors.New("message is required")
    }
    return nil
}

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	UserID uint `gorm:"not null" json:"user_id" form:"user_id" constraint:"OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo ID is required"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil

	return err
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := ValidateMessage(c.Message); err != nil {
        return err
    }

	err = nil
	return
}