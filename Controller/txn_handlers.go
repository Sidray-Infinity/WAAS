package Controller

import (
	"net/http"
	"waas/Domain"
)

type TransactionHandler struct {
	transactionDomain Domain.TransactionDomain
}

func (t *TransactionHandler) GenerateCSV(w http.ResponseWriter, r *http.Request) {
	t.transactionDomain.GenerateCSV()
}
