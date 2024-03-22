package models

import (
	"MyGram/helpers"
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func ValidateAge(age uint8) error {
    if age < 8 {
        return errors.New("Age must be 8 or greater")
    }
    return nil
}

func ValidateEmail(email string) error {
    if email == "" {
        return errors.New("Email is required")
    }
    return nil
}

func ValidateUsername(username string) error {
    if username == "" {
        return errors.New("Username is required")
    }
    return nil
}

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a mininum length of 6 characters"`
	Age			 uint8 			`gorm:"not null" json:"age" form:"age"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Socialmedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialMedias"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if err := ValidateAge(u.Age); err != nil {
        return err
    }

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := ValidateAge(u.Age); err != nil {
        return err
    }

	if err := ValidateEmail(u.Email); err != nil {
        return err
    }
	
	if err := ValidateUsername(u.Username); err != nil {
        return err
    }

	err = nil
	return
}
