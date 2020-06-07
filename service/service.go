package service

import (
	"github.com/caioeverest/transactions-api/logger"
	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/repository"
)

var (
	log *logger.Logger
)

func Start(operationsRepo repository.Interface, accountsRepo repository.Interface, transactionsRepo repository.Interface) (acc *AccountsService, trx *TransactionService) {
	log = logger.Get()
	op := &OperationsService{operationsRepo}
	acc = &AccountsService{accountsRepo}
	trx = &TransactionService{transactionsRepo, acc, op}
	return
}

type OperationsInterface interface {
	RecoverDescription(id int) (string, error)
}

type AccountsInterface interface {
	Create(document string) (*model.Account, error)
	FindAll() ([]model.Account, error)
	FindById(accountId int) (*model.Account, error)
}

type TransactionsInterface interface {
	Create(accountsId int, operationTypeId int, amount float64) (*TransactionDTO, error)
	FindAll() ([]TransactionDTO, error)
}
