package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/service"
	"github.com/caioeverest/transactions-api/util"
	"github.com/stretchr/testify/assert"
)

var (
	createAccountWithEmptyBody          = AccountCreationReq{}
	createAccountWithNomNumericDocument = AccountCreationReq{Document: NomNumericDoc}
	createAccountWithConflictBody       = AccountCreationReq{Document: ExtDocument}
	createAccountBody                   = AccountCreationReq{Document: NewDocument}
)

func TestAccountsHandler_CreateShouldFailWhenNotSpecifyingContentType(t *testing.T) {
	beforeEach()
	router.Post("/accounts", accTarget.Create)
	bodyMarshal, _ := json.Marshal(createAccountWithEmptyBody)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewReader(bodyMarshal))
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_CreateShouldFailWhenRequestComesWithEmptyBody(t *testing.T) {
	beforeEach()
	router.Post("/accounts", accTarget.Create)
	req := mockReq("POST", "/accounts", createAccountWithEmptyBody)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_CreateShouldFailWhenGivenDocumentHaveCharacters(t *testing.T) {
	beforeEach()
	router.Post("/accounts", accTarget.Create)
	req := mockReq("POST", "/accounts", createAccountWithNomNumericDocument)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_CreateShouldFailWhenDocumentAreAlreadyRegister(t *testing.T) {
	beforeEach()
	accMockServ.On("Create", ExtDocument).Return(service.DocumentConflictError).Once()
	router.Post("/accounts", accTarget.Create)
	req := mockReq("POST", "/accounts", createAccountWithConflictBody)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusConflict, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_CreateShouldFailWhenHaveUnexpectedError(t *testing.T) {
	beforeEach()
	accMockServ.On("Create", NewDocument).Return(util.Error("test")).Once()
	router.Post("/accounts", accTarget.Create)
	req := mockReq("POST", "/accounts", createAccountBody)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_CreateWithSuccess(t *testing.T) {
	beforeEach()
	accMockServ.On("Create", NewDocument).Return(&model.Account{}, nil).Once()
	router.Post("/accounts", accTarget.Create)
	req := mockReq("POST", "/accounts", createAccountBody)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_FindByIdShouldFailWhenAccountIDParamIsNotANumber(t *testing.T) {
	beforeEach()
	router.Get("/accounts/:accountID", accTarget.FindById)
	req := mockReq("GET", fmt.Sprintf("/accounts/%s", UnexpectedAccountId), nil)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_FindByIdShouldFailWhenAccountNotFound(t *testing.T) {
	beforeEach()
	accMockServ.On("FindById", NotExtAccountId).Return(util.Error("not found")).Once()
	router.Get("/accounts/:accountID", accTarget.FindById)
	req := mockReq("GET", fmt.Sprintf("/accounts/%d", NotExtAccountId), nil)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_FindByIdWithSuccess(t *testing.T) {
	beforeEach()
	accMockServ.On("FindById", AccountId).Return(&model.Account{}, nil).Once()
	router.Get("/accounts/:accountID", accTarget.FindById)
	req := mockReq("GET", fmt.Sprintf("/accounts/%d", AccountId), nil)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_FindAllShouldFailWhenHaveUnexpectedError(t *testing.T) {
	beforeEach()
	accMockServ.On("FindAll").Return(util.Error("test")).Once()
	router.Get("/accounts", accTarget.FindAll)
	req := mockReq("GET", "/accounts", nil)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Nil(t, err)
}

func TestAccountsHandler_FindAllWithSuccess(t *testing.T) {
	beforeEach()
	accMockServ.On("FindAll").Return([]model.Account{}, nil).Once()
	router.Get("/accounts", accTarget.FindAll)
	req := mockReq("GET", "/accounts", nil)
	res, err := router.Test(req)
	assert.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Nil(t, err)
}
