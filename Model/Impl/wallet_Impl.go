package Impl

import (
	"errors"
	"log"
	"strconv"
	"time"
	entity "waas/Model/entity"
	"waas/Model/view"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func getBalanceRedis(walletId int) (float64, bool) {
	vals, err := rdb.Get(ctx, strconv.Itoa(walletId)).Result()
	if err != nil {
		log.Println("Cannot fetch from catche:", err)
		return -1, false
	}
	if len(vals) > 0 {
		val, err := strconv.ParseFloat(vals, 64)
		if err != nil {
			log.Println("Redis: Cannot convert to float", err)
			return -1, false
		}
		return val, true
	}
	return -1, false
}

func setBalanceRedis(walletId int, balance float64, expiry time.Duration) {
	_, err = rdb.Set(ctx, strconv.Itoa(walletId),
		strconv.FormatFloat(balance, 'f', 6, 64), expiry).Result()

	if err != nil {
		log.Println("Cannot set on cache:", err)
	}
}

func GetBalance(walletId int) (*float64, error) {
	val, found := getBalanceRedis(walletId)
	if found {
		return &val, nil
	}
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

	setBalanceRedis(walletId, wallet.Balance, 0)

	return &wallet.Balance, nil
}

func GetStatus(walletId int) (*bool, error) {
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

	return &wallet.IsBlocked, nil
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
		log.Println("Cannot start DB transaction:", err)
		return -1, -1, err
	}

	var transaction entity.Transaction
	var wallet entity.Wallet
	var isFailedTransaction bool = false

	// err = tx.Set("gorm:query_option", "FOR UPDATE").First(&wallet, walletId).Error
	err = tx.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&wallet, walletId).Error

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

	transaction.Amount = updateReq.UpdateAmount
	transaction.Type = true
	transaction.Wallet = wallet
	transaction.Time = time.Now()
	transaction.Status = "pending"

	err = tx.Create(&transaction).Error
	if err != nil {
		log.Println("Cannot create transaction:", err)
		tx.Rollback()
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
		transaction.Status = "failed"
		isFailedTransaction = true
	} else {
		transaction.Status = "completed"
	}

	err = tx.Save(&transaction).Error
	if err != nil {
		log.Println("Cannot update transaction:", err)
		tx.Rollback()
		return -1, -1, err
	}

	if isFailedTransaction {
		return -1, -1, err
	}

	time.Sleep(10 * time.Second)
	log.Println("---------------------------------DONE")

	if err := tx.Commit().Error; err != nil {
		log.Println("Cannot commit transaction:", err)
		tx.Rollback()
		return -1, -1, err
	}
	return wallet.Balance, transaction.ID, nil

}

func WalletStatus(updateReq *view.StatusUpdate, walletId int) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		log.Println("Cannot start DB transaction:", err)
		return err
	}
	var wallet entity.Wallet
	err = tx.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&wallet, walletId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Record not found for Wallet ID:", walletId)
		return err
	}
	if err != nil {
		log.Println("Cannot fetch wallet", err)
		return err
	}

	wallet.IsBlocked = updateReq.NewStatus
	err = tx.Save(&wallet).Error
	if err != nil {
		log.Println("Cannot update wallet status:", err)
		return err
	}
	// time.Sleep(10 * time.Second)
	// log.Println("---------------------------------DONE STATUS")
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
