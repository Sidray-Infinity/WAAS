package Model

import (
	entity "waas/Model/entity"
	"waas/Model/view"
)

type WalletModel interface {
	GetWallet(walletId int) (*entity.Wallet, error)
	GetBalance(walletId int) (*float64, error)
	GetStatus(walletId int) (*bool, error)
	RegisterWallet(newWallet *entity.Wallet) error
	WalletBalance(updateReq *view.BalanceUpdate, walletId int) (float64, int, error)
	WalletStatus(updateReq *view.StatusUpdate, walletId int) error
}
