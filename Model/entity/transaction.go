package Model

import "time"

type Transaction struct {
	ID       int       `json:"transaction_id"`
	Wallet   Wallet    `gorm:"foreignKey:WalletId"`
	WalletId int       `json:"fk_wallet_id"`
	Amount   float64   `json:"amount"`
	Type     bool      `json:"type"`
	Time     time.Time `json:"time"`
	Status   string    `json:"status"`
}
