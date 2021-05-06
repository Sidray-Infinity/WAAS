package Impl

import (
	"errors"
	"log"
	"sync"
	"time"
	entity "waas/Model/entity"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func AddWallet(userId int) {
	db, _ := gorm.Open("mysql", address)
	var user entity.User

	err = db.First(&user, userId).Error
	if err != nil {
		log.Println("Invalid User ID! ", err)
		return
	}

	var wallet entity.Wallet
	wallet.User = user
	err = db.Create(&wallet).Error
	if err != nil {
		log.Println("Error while adding wallet:", err)
	}
	db.DB().Close()

}

func Credit(walletId int, amount float64) {
	db, _ := gorm.Open("mysql", address)

	var wallet entity.Wallet
	var transaction entity.Transaction

	err = db.First(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Invalid Wallet Id")
		return
	}

	if wallet.IsBlocked {
		log.Println("Cannot transact on blocked wallets")
		return
	}
	if walletMutx[walletId] == nil {
		walletMutx[walletId] = &sync.Mutex{}
	}

	walletMutx[walletId].Lock()
	log.Println("CREDIT MUTEX:", walletId)
	wallet.Balance += amount
	txTime := time.Now()
	db.Save(&wallet)
	// log.Println("---------------------------------SLEEPING", walletId)
	time.Sleep(time.Second * 10)
	// log.Println("---------------------------------DONE", walletId)
	log.Println("DONE:", walletId)
	walletMutx[walletId].Unlock()

	transaction.Amount = amount
	transaction.Type = true
	transaction.Wallet = wallet
	transaction.Time = txTime

	db.Create(&transaction)

	db.DB().Close()
}

func Debit(walletId int, amount float64) {
	db, _ := gorm.Open("mysql", address)
	var wallet entity.Wallet
	var transaction entity.Transaction

	err = db.First(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Invalid Wallet Id")
		return
	}

	if wallet.IsBlocked {
		log.Println("Cannot transact on blocked wallets")
		return
	}

	if wallet.Balance-amount < 0 {
		log.Println("Balance too low for the given amount")
		return
	}

	if walletMutx[walletId] == nil {
		walletMutx[walletId] = &sync.Mutex{}
	}

	walletMutx[walletId].Lock()
	log.Println("DEBIT MUTEX:", walletId)

	wallet.Balance -= amount
	txTime := time.Now()

	db.Save(&wallet)
	walletMutx[walletId].Unlock()

	transaction.Amount = amount
	transaction.Type = false
	transaction.Wallet = wallet
	transaction.Time = txTime

	db.Create(&transaction)

	db.DB().Close()
}

func Block(walletId int) {
	db, _ := gorm.Open("mysql", address)

	var wallet entity.Wallet
	err = db.First(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Invalid Wallet Id")
		return
	}

	wallet.IsBlocked = true
	db.Save(&wallet)
	db.DB().Close()
}

func UnBlock(walletId int) {
	db, _ := gorm.Open("mysql", address)

	var wallet entity.Wallet
	err = db.First(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Invalid Wallet Id")
		return
	}

	wallet.IsBlocked = false
	db.Save(&wallet)
	db.DB().Close()
}
