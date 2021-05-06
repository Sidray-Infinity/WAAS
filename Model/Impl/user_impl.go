package Impl

import (
	"log"
	entity "waas/Model/entity"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Get() []entity.User {
	// TEST APIs
	db, _ := gorm.Open("mysql", address)
	var users []entity.User
	db.Find(&users)
	db.DB().Close()
	return users
}

func Register(newUser entity.User) {
	db, _ := gorm.Open("mysql", address)
	err = db.Create(&newUser).Error
	if err != nil {
		log.Println("Error while registering user:", err)
	}
	db.DB().Close()
}
