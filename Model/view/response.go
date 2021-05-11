package view

type BalanceUpdateResp struct {
	TransactionId  int     `json:"transaction_id"`
	UpdatedBalance float64 `json:"updated_balance"`
}
