package Domain

import (
	"waas/Model/Impl"
)

func AddWallet(userId int) {
	Impl.AddWallet(userId)
}

func Credit(walletId int, amount float64) {
	Impl.Credit(walletId, amount)
}

func Debit(walletId int, amount float64) {
	Impl.Debit(walletId, amount)
}

func Block(walletId int) {
	Impl.Block(walletId)
}

func UnBlock(walletId int) {
	Impl.UnBlock(walletId)
}
