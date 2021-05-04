package manager

import (
	"waas/dao"
	"waas/model"
)

func Get() []model.User {
	return dao.Get()
}

func Register(newUser model.User) {
	dao.Register(newUser)
}

func AddWallet(userId int) {
	dao.AddWallet(userId)
}

func Credit(walletId int, amount float64) {
	dao.Credit(walletId, amount)
}

func Debit(walletId int, amount float64) {
	dao.Debit(walletId, amount)
}
