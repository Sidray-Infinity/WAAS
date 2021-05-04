package model

type Transaction struct {
	ID       int  `json:"transaction_id"`
	User     User `json:"fk_user_id" gorm:"foreignKey:UserId"`
	UserId   int
	Wallet   Wallet `json:"fk_wallet_id" gorm:"foreignKey:WalletId"`
	WalletId int
	Amount   float64 `json:"amount"`
	Type     bool    `json:"type"`
}
