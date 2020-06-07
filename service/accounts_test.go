package service

import (
	"testing"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/util"
	"github.com/stretchr/testify/assert"
)

func TestAccountsService_CreateShouldFailWhenDocumentAlreadyInUse(t *testing.T) {
	beforeEach()
	acc := &model.Account{DocumentNumber: ExtDocument}
	repoMock.On("Find", acc).Return(nil)
	out, err := accTarget.Create(ExtDocument)
	assert.Equal(t, DocumentConflictError, err)
	assert.Nil(t, out)
}

func TestAccountsService_CreateShouldFailWhenRepositoryReturnError(t *testing.T) {
	beforeEach()
	e := util.Error("test")
	acc := &model.Account{DocumentNumber: NewDocument}
	repoMock.On("Find", acc).Return(util.Error("not found"))
	repoMock.On("Save", acc).Return(e)
	out, err := accTarget.Create(NewDocument)
	assert.Equal(t, e, err)
	assert.Nil(t, out)
}

func TestAccountsService_CreateShouldSucceed(t *testing.T) {
	beforeEach()
	acc := &model.Account{DocumentNumber: NewDocument}
	repoMock.On("Find", acc).Return(util.Error("not found"))
	repoMock.On("Save", acc).Return(nil)
	out, err := accTarget.Create(NewDocument)
	assert.Nil(t, err)
	assert.Equal(t, NewDocument, out.DocumentNumber)
}

func TestAccountsService_FindAllShouldFailWhenRepositoryReturnError(t *testing.T) {
	beforeEach()
	e := util.Error("test")
	accounts := make([]model.Account, 0)
	repoMock.On("FindAll", &accounts).Return(e)
	out, err := accTarget.FindAll()
	assert.Equal(t, e, err)
	assert.Nil(t, out)
}

func TestAccountsService_FindAllShouldSucceed(t *testing.T) {
	beforeEach()
	accounts := make([]model.Account, 0)
	repoMock.On("FindAll", &accounts).Return(nil)
	out, err := accTarget.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, accounts, out)
}

func TestAccountsService_FindByIdShouldFailWhenRepositoryReturnError(t *testing.T) {
	beforeEach()
	e := util.Error("test")
	account := model.Account{}
	repoMock.On("FindById", NotExtAccountId, &account).Return(e)
	out, err := accTarget.FindById(NotExtAccountId)
	assert.Equal(t, e, err)
	assert.Nil(t, out)
}

func TestAccountsService_FindByIdShouldSucceed(t *testing.T) {
	beforeEach()
	account := model.Account{}
	repoMock.On("FindById", ExtAccountId, &account).Return(nil)
	out, err := accTarget.FindById(ExtAccountId)
	assert.Nil(t, err)
	assert.Equal(t, &account, out)
}
