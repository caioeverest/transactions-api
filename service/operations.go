package service

import (
	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/repository"
)

type OperationsService struct {
	repo repository.Interface
}

func (s *OperationsService) RecoverDescription(id int) (desc string, err error) {
	operation := &model.Operation{}
	if err = s.repo.FindById(int(id), operation); err != nil {
		return
	}
	return operation.Description, err
}
