package Impl

import (
	"errors"
	"log"
	entity "waas/Model/entity"

	"gorm.io/gorm"
)

type UserModelImpl struct{}

func (u *UserModelImpl) GetUser(userId int) (*entity.User, error) {
	var user entity.User
	err = db.Find(&user, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Record not found for User ID:", userId)
		return nil, err
	}
	if err != nil {
		log.Println("Cannot fetch user", err)
		return nil, err
	}

	return &user, nil
}

func (u *UserModelImpl) RegisterUser(newUser entity.User) error {
	err = db.Create(&newUser).Error
	if err != nil {
		log.Println("Eror while registering user:", err)
	}
	return err
}

func (u *UserModelImpl) ValidateUsername(newUser entity.User) bool {
	err = db.Where("user_name = ?", newUser.UserName).First(&newUser).Error
	if err == gorm.ErrRecordNotFound {
		return true
	} else if err != nil {
		log.Println("Cannot validate username:", err)
	}
	return false
}

func (u *UserModelImpl) ValidateKYC(newUser entity.User) bool {
	err = db.Where("aadhar_number = ?", newUser.AadharNumber).First(&newUser).Error
	if err == gorm.ErrRecordNotFound {
		return true
	} else if err != nil {
		log.Println("Cannot validate Aadhar Number:", err)
	}
	return false
}
