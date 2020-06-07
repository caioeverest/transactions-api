package service

import (
	"sync"
	"time"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/repository"
)

const (
	COMPRA_A_VISTA = iota + 1
	COMPRA_PARCELADA
	SAQUE
	PAGAMENTO
)

type TransactionService struct {
	repo    repository.Interface
	accServ AccountsInterface
	opServ  OperationsInterface
}

type TransactionDTO struct {
	TransactionID int       `json:"transaction_id"`
	AccountID     int       `json:"account_id"`
	Operation     string    `json:"operation"`
	Amount        float64   `json:"amount"`
	EventDate     time.Time `json:"event_date"`
}

//Create a new transaction
func (s *TransactionService) Create(accountsId int, operationTypeId int, amount float64) (trx *TransactionDTO, err error) {
	if err = s.transactionValidations(accountsId, operationTypeId, amount); err != nil {
		return
	}

	tmp := &model.Transaction{
		AccountID:       accountsId,
		OperationTypeId: operationTypeId,
		Amount:          calculateAmount(operationTypeId, amount),
	}

	log.Info("New request for transaction creation")
	if err = s.repo.Save(tmp); err != nil {
		log.Errorf("Unexpected error creating transaction - ERROR %+v", err)
		return nil, err
	}
	log.Infof("Transaction %d created", tmp.TransactionID)

	trx = s.converterToDTO(tmp)

	return trx, nil
}

//Recover all transactions from repository
func (s TransactionService) FindAll() ([]TransactionDTO, error) {
	var (
		transactions    = make([]model.Transaction, 0)
		transactionsDTO = make([]TransactionDTO, 0)
		wg              sync.WaitGroup
	)
	if err := s.repo.FindAll(&transactions); err != nil {
		log.Errorf("Unexpected error listing transactions - ERROR %+v", err)
		return nil, err
	}

	wg.Add(len(transactions))
	go func() {
		for _, t := range transactions {
			transactionsDTO = append(transactionsDTO, *s.converterToDTO(&t))
			wg.Done()
		}
	}()
	wg.Wait()

	return transactionsDTO, nil
}

func (s TransactionService) converterToDTO(src *model.Transaction) *TransactionDTO {
	var OperationDescription string
	OperationDescription, _ = s.opServ.RecoverDescription(src.OperationTypeId)
	return &TransactionDTO{
		TransactionID: src.TransactionID,
		AccountID:     src.AccountID,
		Operation:     OperationDescription,
		Amount:        src.Amount,
		EventDate:     src.EventDate,
	}
}

func (s TransactionService) transactionValidations(accountsId int, operationTypeId int, amount float64) (err error) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		if amount < 0 {
			err = AmountCantBeNegativeError
		}
		wg.Done()
	}()
	go func() {
		if _, e := s.accServ.FindById(accountsId); e != nil {
			err = AccountNotFoundError
		}
		wg.Done()
	}()

	go func() {
		if _, e := s.opServ.RecoverDescription(operationTypeId); e != nil {
			err = OperationNotFoundError
		}
		wg.Done()
	}()
	wg.Wait()
	return
}

func calculateAmount(operationTypeId int, amount float64) float64 {
	if operationTypeId != PAGAMENTO {
		return amount * -1
	}
	return amount
}
