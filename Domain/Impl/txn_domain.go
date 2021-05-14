package Domain

import (
	"waas/Model"
	ModelImpl "waas/Model/Impl"
)

type TransactionDomainImpl struct {
	transactionModel Model.TransactionModel
}

func (t *TransactionDomainImpl) GenerateCSV() {
	t.transactionModel = &ModelImpl.TransactionModelImpl{}
	t.transactionModel.GenerateCSV()
}
