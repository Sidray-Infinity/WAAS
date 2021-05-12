package Domain

import (
	"waas/Model"
)

type TransactionDomainImpl struct {
	transactionModel Model.TransactionModel
}

func (t *TransactionDomainImpl) GenerateCSV() {
	t.transactionModel.GenerateCSV()
}
