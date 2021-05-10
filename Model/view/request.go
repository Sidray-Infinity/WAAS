package view

type BalanceUpdate struct {
	UpdateAmount float64 `json:"amount"`
	UpdateType   bool    `json:"type"`
}

type StatusUpdate struct {
	NewStatus bool `json:"status"`
}
