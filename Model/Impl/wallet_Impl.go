package Impl

import (
	"errors"
	"log"
	"time"
	entity "waas/Model/entity"
	"waas/Model/view"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

func GetWallet(walletId int) (*entity.Wallet, error) {
	wallet := &entity.Wallet{}
	err = db.Find(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Record not found for Wallet ID:", walletId)
		return nil, err
	}
	if err != nil {
		log.Println("Cannot fetch wallet", err)
		return nil, err
	}

	return wallet, nil
}

func RegisterWallet(newWallet *entity.Wallet) error {
	user, err := GetUser(newWallet.UserId)
	if err != nil {
		return err
	}

	newWallet.User = *user
	err = db.Create(&newWallet).Error
	if err != nil {
		log.Println("Eror while registering wallet:", err)
	}
	return err
}

func WalletBalance(updateReq *view.BalanceUpdate, walletId int) (float64, int, error) {

	var transaction entity.Transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return -1, -1, err
	}

	wallet, err := GetWallet(walletId)
	if err != nil {
		return -1, -1, nil
	}

	if err := tx.Set("gorm:query_option", "FOR UPDATE").Error; err != nil {
		tx.Rollback()
		return -1, -1, err
	}

	if wallet.IsBlocked {
		log.Println("Cannot transact on blocked wallets")
		return -1, -1, err
	}

	if updateReq.UpdateType {
		// Credit
		transaction.Type = true
		wallet.Balance += updateReq.UpdateAmount
	} else {
		// Debit
		transaction.Type = false
		wallet.Balance -= updateReq.UpdateAmount
	}

	err = db.Save(&wallet).Error
	if err != nil {
		log.Println("Cannot update wallet balance:", err)
		return -1, -1, err
	}

	txTime := time.Now()
	transaction.Amount = updateReq.UpdateAmount
	transaction.Type = true
	transaction.Wallet = *wallet
	transaction.Time = txTime

	err = db.Create(&transaction).Error
	if err != nil {
		log.Println("Cannot create transaction:", err)
		return -1, -1, err
	}
	return wallet.Balance, transaction.ID, nil

}

func WalletStatus(updateReq *view.StatusUpdate, walletId int) error {

	wallet, err := GetWallet(walletId)
	if err != nil {
		return err
	}
	wallet.IsBlocked = updateReq.NewStatus
	db.Save(&wallet)
	return nil
}
