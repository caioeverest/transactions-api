package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/caioeverest/transactions-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

const (
	AccountId           = 1
	UnexpectedAccountId = "1231nm"
	NotExtAccountId     = 789137
	Amount              = 1384.32
	UnexpectedOP        = 1237
	NewDocument         = "3781278379812"
	ExtDocument         = "3781278379812"
	NomNumericDoc       = "asda908129nk"
)

var (
	router *fiber.App

	accMockServ *service.AccountsMock
	accTarget   AccountsHandler

	trxMockServ *service.TransactionsMock
	trxTarget   TransactionsHandler
	validate    *validator.Validate
)

func beforeEach() {
	validate = validator.New()
	accMockServ = &service.AccountsMock{}
	accTarget = AccountsHandler{accMockServ, validate}

	trxMockServ = &service.TransactionsMock{}
	trxTarget = TransactionsHandler{trxMockServ, validate}

	router = fiber.New()
}

func mockReq(method, path string, body interface{}) *http.Request {
	bodyMarshal, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, path, bytes.NewReader(bodyMarshal))
	req.Header.Set("Content-Type", "application/json")
	return req
}
