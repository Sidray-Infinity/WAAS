package Controller

import (
	"net/http"
	"waas/Domain"
	DomainImpl "waas/Domain/Impl"
)

type TransactionHandler struct {
	transactionDomain Domain.TransactionDomain
}

func (t *TransactionHandler) GenerateCSV(w http.ResponseWriter, r *http.Request) {
	t.transactionDomain = &DomainImpl.TransactionDomainImpl{}
	t.transactionDomain.GenerateCSV()
}
