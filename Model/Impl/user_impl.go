package Impl

import (
	"errors"
	"log"
	entity "waas/Model/entity"

	"gorm.io/gorm"
)

func GetUser(userId int) (*entity.User, error) {
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

func RegisterUser(newUser entity.User) error {
	err = db.Create(&newUser).Error
	if err != nil {
		log.Println("Eror while registering user:", err)
	}
	return err
}
