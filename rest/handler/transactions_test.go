package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/caioeverest/transactions-api/service"
	"github.com/caioeverest/transactions-api/util"
	"github.com/stretchr/testify/assert"
)

var (
	createPaymentTransactionBody = TransactionCreationReq{
		AccountID:       AccountId,
		OperationTypeId: service.PAGAMENTO,
		Amount:          Amount,
	}

	createFailTransactionBodyAbsentParam = TransactionCreationReq{
		AccountID:       AccountId,
		OperationTypeId: service.PAGAMENTO,
	}

	createFailTransactionBodyFailNonexistentAccount = TransactionCreationReq{
		AccountID:       NotExtAccountId,
		OperationTypeId: service.PAGAMENTO,
		Amount:          Amount,
	}

	createFailTransactionBodyUnexpectedOperation = TransactionCreationReq{
		AccountID:       AccountId,
		OperationTypeId: UnexpectedOP,
		Amount:          Amount,
	}

	createFailTransactionBodyNegativeAmount = TransactionCreationReq{
		AccountID:       AccountId,
		OperationTypeId: service.PAGAMENTO,
		Amount:          -Amount,
	}
)

func TestTransactionsHandler_CreateShouldFailNotSpecifyingContentType(t *testing.T) {
	beforeEach()
	router.Post("/transactions", trxTarget.Create)
	bodyMarshal, _ := json.Marshal(createFailTransactionBodyAbsentParam)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewReader(bodyMarshal))
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateShouldFailWhenDontHaveBody(t *testing.T) {
	beforeEach()
	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createFailTransactionBodyAbsentParam)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateShouldFailWhenAmountIsNegative(t *testing.T) {
	beforeEach()
	trxMockServ.On("Create", AccountId, service.PAGAMENTO, -Amount).
		Return(service.AmountCantBeNegativeError).Once()

	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createFailTransactionBodyNegativeAmount)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateShouldFailWhenAccountNotFound(t *testing.T) {
	beforeEach()
	trxMockServ.On("Create", NotExtAccountId, service.PAGAMENTO, Amount).
		Return(service.AccountNotFoundError).Once()

	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createFailTransactionBodyFailNonexistentAccount)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateShouldFailWhenUseUnexpectedOperation(t *testing.T) {
	beforeEach()
	trxMockServ.On("Create", AccountId, UnexpectedOP, Amount).
		Return(service.OperationNotFoundError).Once()

	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createFailTransactionBodyUnexpectedOperation)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateShouldFailWhenHaveAnUnexpectedError(t *testing.T) {
	beforeEach()
	trxMockServ.On("Create", AccountId, service.PAGAMENTO, Amount).
		Return(util.Error("test")).Once()

	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createPaymentTransactionBody)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_CreateWithSuccess(t *testing.T) {
	beforeEach()

	trxMockServ.On("Create", AccountId, service.PAGAMENTO, Amount).Return(&service.TransactionDTO{}, nil).Once()
	router.Post("/transactions", trxTarget.Create)
	req := mockReq("POST", "/transactions", createPaymentTransactionBody)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_FindAllShouldFailWhenServiceReturnError(t *testing.T) {
	beforeEach()
	trxMockServ.On("FindAll").Return(util.Error("test")).Once()

	router.Get("/transactions", trxTarget.FindAll)
	req := mockReq("GET", "/transactions", nil)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Nil(t, err)
}

func TestTransactionsHandler_FindAllWithSuccess(t *testing.T) {
	beforeEach()
	trxMockServ.On("FindAll").Return([]service.TransactionDTO{}, nil).Once()

	router.Get("/transactions", trxTarget.FindAll)
	req := mockReq("GET", "/transactions", nil)
	res, err := router.Test(req)

	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}
