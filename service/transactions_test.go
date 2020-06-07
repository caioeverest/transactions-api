package service

import (
	"testing"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/util"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_CreateShouldReturnErrorWhenHaveNegativeAmount(t *testing.T) {
	beforeEach()
	accMock.On("FindById", ExtAccountId).Return(nil)
	opMock.On("RecoverDescription", COMPRA_A_VISTA).Return("", nil)

	dto, err := trxTarget.Create(ExtAccountId, COMPRA_A_VISTA, -Amount)
	assert.Equal(t, AmountCantBeNegativeError, err)
	assert.Nil(t, dto)
}

func TestTransactionService_CreateShouldReturnErrorAccountDoesNotExits(t *testing.T) {
	beforeEach()
	accMock.On("FindById", NotExtAccountId).Return(util.Error("error"))
	opMock.On("RecoverDescription", COMPRA_A_VISTA).Return("test", nil)

	dto, err := trxTarget.Create(NotExtAccountId, COMPRA_A_VISTA, Amount)
	assert.Equal(t, AccountNotFoundError, err)
	assert.Nil(t, dto)
}

func TestTransactionService_CreateShouldReturnErrorWhenOperationDoesNotExist(t *testing.T) {
	beforeEach()
	accMock.On("FindById", ExtAccountId).Return(nil)
	opMock.On("RecoverDescription", 7890).Return("", util.Error("error"))

	dto, err := trxTarget.Create(ExtAccountId, 7890, Amount)
	assert.Equal(t, OperationNotFoundError, err)
	assert.Nil(t, dto)
}

func TestTransactionService_CreateShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	beforeEach()

	expectedTRX := &model.Transaction{
		AccountID:       ExtAccountId,
		OperationTypeId: COMPRA_A_VISTA,
		Amount:          -Amount,
	}
	e := util.Error("error")

	accMock.On("FindById", ExtAccountId).Return(nil)
	opMock.On("RecoverDescription", COMPRA_A_VISTA).Return("test", nil)
	repoMock.On("Save", expectedTRX).Return(e)

	dto, err := trxTarget.Create(ExtAccountId, COMPRA_A_VISTA, Amount)
	assert.Equal(t, e, err)
	assert.Nil(t, dto)
}

func TestTransactionService_CreateShouldReturnSuccessWhenPassAllValidations(t *testing.T) {
	beforeEach()

	expectedTRX := &model.Transaction{
		AccountID:       ExtAccountId,
		OperationTypeId: PAGAMENTO,
		Amount:          Amount,
	}

	accMock.On("FindById", ExtAccountId).Return(nil)
	opMock.On("RecoverDescription", PAGAMENTO).Return("test", nil)
	repoMock.On("Save", expectedTRX).Return(nil)

	dto, err := trxTarget.Create(ExtAccountId, PAGAMENTO, Amount)
	assert.Nil(t, err)
	assert.NotNil(t, dto)
}

func TestTransactionService_FindAllShouldReturnErrorWhenRepositoryReturnError(t *testing.T) {
	beforeEach()
	e := util.Error("error")
	transactions := make([]model.Transaction, 0)
	repoMock.On("FindAll", &transactions).Return(e)
	dto, err := trxTarget.FindAll()
	assert.Equal(t, e, err)
	assert.Nil(t, dto)
}

func TestTransactionService_FindAllWithSuccess(t *testing.T) {
	beforeEach()
	transactions := make([]model.Transaction, 0)
	repoMock.On("FindAll", &transactions).Return(nil)
	dto, err := trxTarget.FindAll()
	assert.Nil(t, err)
	assert.Len(t, dto, 0)
}
