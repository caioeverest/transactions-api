package service

import (
	"github.com/caioeverest/transactions-api/logger"
	"github.com/caioeverest/transactions-api/repository"
)

const (
	ExtAccountId    = 1234
	NotExtAccountId = 4321
	Amount          = 123.45
	NewDocument     = "877897834"
	ExtDocument     = "281317823"
)

var (
	repoMock *repository.Mock
	opMock   *OperationsMock
	accMock  *AccountsMock

	opTarget  *OperationsService
	accTarget *AccountsService
	trxTarget *TransactionService
)

func beforeEach() {
	logger.Start()
	log = logger.Get()

	repoMock = &repository.Mock{}
	opMock = &OperationsMock{}
	accMock = &AccountsMock{}

	opTarget = &OperationsService{repoMock}
	accTarget = &AccountsService{repoMock}
	trxTarget = &TransactionService{repoMock, accMock, opMock}
}
