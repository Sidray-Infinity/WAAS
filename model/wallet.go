package model

type Wallet struct {
	ID        int     `json:"wallet_id"`
	User      User    `gorm:"foreignKey:UserId"`
	UserId    int     `json:"user_id"`
	Balance   float64 `json:"balance"`
	IsBlocked int     `json:"is_blocked"`
}
