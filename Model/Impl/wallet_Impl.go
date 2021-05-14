package Impl

import (
	"errors"
	"log"
	"time"
	"waas/Model"
	entity "waas/Model/entity"
	"waas/Model/view"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletModelImpl struct {
	userModel Model.UserModel
}

func (w *WalletModelImpl) GetWallet(walletId int) (*entity.Wallet, error) {
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

func (w *WalletModelImpl) GetBalance(walletId int) (*float64, error) {
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

func (w *WalletModelImpl) GetStatus(walletId int) (*bool, error) {
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

func (w *WalletModelImpl) RegisterWallet(newWallet *entity.Wallet) error {
	w.userModel = &UserModelImpl{}
	user, err := w.userModel.GetUser(newWallet.UserId)
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

func (w *WalletModelImpl) WalletBalance(updateReq *view.BalanceUpdate, walletId int) (float64, int, error) {

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

	// Apply Shared lock on the DB row, if found
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

	//deleteBalanceRedis(walletId) // Deleting redis entry for data consistency
	/*
		========================================================
		WHAT IF CONTEXT SWITCHES HERE?
	*/
	// for i := 0; i < 15; i++ {
	// 	log.Printf("%d ", i)
	// 	time.Sleep(time.Second)
	// }
	if err := tx.Commit().Error; err != nil {
		log.Println("Cannot commit transaction:", err)
		tx.Rollback()
		return -1, -1, err
	}
	// for i := 0; i < 15; i++ {
	// 	log.Printf("%d ", i)
	// 	time.Sleep(time.Second)
	// }

	// context switch : getBalance
	setBalanceRedis(walletId, wallet.Balance, 0) // Update cache with new balance

	return wallet.Balance, transaction.ID, nil

}

func (w *WalletModelImpl) WalletStatus(updateReq *view.StatusUpdate, walletId int) error {
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
