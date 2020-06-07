package service

import (
	"github.com/caioeverest/transactions-api/model"
	"github.com/stretchr/testify/mock"
)

type OperationsMock struct {
	mock.Mock
}

func (m *OperationsMock) RecoverDescription(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

type AccountsMock struct {
	mock.Mock
}

func (m *AccountsMock) Create(document string) (*model.Account, error) {
	args := m.Called(document)
	acc, ok := args.Get(0).(*model.Account)
	if !ok {
		return nil, args.Error(0)
	}
	return acc, args.Error(1)
}
func (m *AccountsMock) FindAll() ([]model.Account, error) {
	args := m.Called()
	acc, ok := args.Get(0).([]model.Account)
	if !ok {
		return nil, args.Error(0)
	}
	return acc, args.Error(1)
}
func (m *AccountsMock) FindById(accountId int) (*model.Account, error) {
	args := m.Called(accountId)
	acc, ok := args.Get(0).(*model.Account)
	if !ok {
		return nil, args.Error(0)
	}
	return acc, args.Error(1)
}

type TransactionsMock struct {
	mock.Mock
}

func (m *TransactionsMock) Create(accountsId int, operationTypeId int, amount float64) (*TransactionDTO, error) {
	args := m.Called(accountsId, operationTypeId, amount)
	trx, ok := args.Get(0).(*TransactionDTO)
	if !ok {
		return nil, args.Error(0)
	}
	return trx, args.Error(1)
}
func (m *TransactionsMock) FindAll() ([]TransactionDTO, error) {
	args := m.Called()
	trx, ok := args.Get(0).([]TransactionDTO)
	if !ok {
		return nil, args.Error(0)
	}
	return trx, args.Error(1)
}
