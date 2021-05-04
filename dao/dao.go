package dao

import (
	"log"
	"math"
	"waas/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var address string = "root:root@tcp(127.0.0.1:3306)/waas?charset=utf8&parseTime=True&loc=Local"

var err error

func Get() []model.User {
	db, _ = gorm.Open("mysql", address)
	var users []model.User
	db.Find(&users)
	db.DB().Close()
	return users
}

func Register(newUser model.User) {
	db, _ = gorm.Open("mysql", address)
	result := db.Create(&newUser)
	log.Println(result) // TODO : Check if the result has failed
	db.DB().Close()
}

func AddWallet(userId int) {
	db, _ = gorm.Open("mysql", address)
	var user model.User

	err = db.First(&user, userId).Error
	if err != nil {
		log.Println("Invalid User ID! ", err)
		return
	}

	var wallet model.Wallet
	wallet.User = user
	result := db.Create(&wallet)
	log.Println(result) // TODO : Check if the result has failed
	db.DB().Close()

}

func Credit(walletId int, amount float64) {
	// TODO : Add condition if wallet is blocked
	db, _ = gorm.Open("mysql", address)
	var wallet model.Wallet
	db.First(&wallet, walletId)
	wallet.Balance += amount
	db.Save(&wallet)
	db.DB().Close()
}

func Debit(walletId int, amount float64) {
	// TODO : Add condition if wallet is blocked
	db, _ = gorm.Open("mysql", address)
	var wallet model.Wallet
	db.First(&wallet, walletId)

	// TODO : Can be improved
	wallet.Balance = math.Max(wallet.Balance-amount, 0)
	db.Save(&wallet)
	db.DB().Close()
}

// func Block(walletId int) {

// }
