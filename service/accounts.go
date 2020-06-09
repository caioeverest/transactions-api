package service

import (
	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/repository"
)

type AccountsService struct {
	repo repository.Interface
}

func (s *AccountsService) Create(document string) (*model.Account, error) {
	var (
		acc = &model.Account{DocumentNumber: document}
		err error
	)

	log.Info("New request for account creation")
	if err = s.repo.Find(acc, acc); err == nil {
		log.Errorf("Document already exist on account %d", acc.AccountID)
		return nil, DocumentConflictError
	}

	if err = s.repo.Save(acc); err != nil {
		log.Errorf("Unexpected error creating account - ERROR %+v", err)
		return nil, err
	}
	log.Infof("Account %d created", acc.AccountID)

	return acc, nil
}

func (s AccountsService) FindAll() ([]model.Account, error) {
	var accounts = make([]model.Account, 0)
	if err := s.repo.FindAll(&accounts); err != nil {
		log.Errorf("Unexpected error listing account - ERROR %+v", err)
		return nil, err
	}
	return accounts, nil
}

func (s AccountsService) FindById(accountId int) (*model.Account, error) {
	account := &model.Account{}
	if err := s.repo.FindById(accountId, account); err != nil {
		log.WithField("account_id", accountId).
			Errorf("Unexpected error searching for account - ERROR %+v", err)
		return nil, err
	}
	return account, nil
}
