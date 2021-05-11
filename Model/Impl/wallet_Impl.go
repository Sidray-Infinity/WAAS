package Impl

import (
	"errors"
	"log"
	"time"
	entity "waas/Model/entity"
	"waas/Model/view"

	"gorm.io/gorm"
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
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println("Cannot start transaction:", err)
		return -1, -1, err
	}

	var transaction entity.Transaction
	var wallet entity.Wallet

	err = tx.Set("gorm:query_option", "FOR UPDATE").First(&wallet, walletId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Record not found for Wallet ID:", walletId)
		return -1, -1, err
	}
	if err != nil {
		log.Println("Cannot fetch wallet", err)
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

	err = tx.Save(&wallet).Error
	if err != nil {
		log.Println("Cannot update wallet balance:", err)
		return -1, -1, err
	}

	txTime := time.Now()
	transaction.Amount = updateReq.UpdateAmount
	transaction.Type = true
	transaction.Wallet = wallet
	transaction.Time = txTime

	err = tx.Create(&transaction).Error
	if err != nil {
		log.Println("Cannot create transaction:", err)
		return -1, -1, err
	}
	time.Sleep(20 * time.Second)
	log.Println("--------------------DONE")
	if err := tx.Commit().Error; err != nil {
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
