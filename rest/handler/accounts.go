package handler

import (
	"net/http"
	"strconv"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

type AccountsHandler struct {
	s        service.AccountsInterface
	validate *validator.Validate
}

type AccountCreationReq struct {
	Document string `json:"document" validate:"required,numeric"`
}

func NewAccountsHandler(s service.AccountsInterface, v *validator.Validate) *AccountsHandler {
	return &AccountsHandler{s, v}
}

// Create godoc
// @Summary Create a new account
// @Description using a document it's possible create a account
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param account body handler.AccountCreationReq true "Create account"
// @Success 200 {object} model.Account
// @Failure 400 {object} handler.JSON
// @Failure 409 {object} handler.JSON
// @Failure 500 {object} handler.JSON
// @Router /accounts [post]
func (h *AccountsHandler) Create(c *fiber.Ctx) {
	var (
		acc  *model.Account
		body = &AccountCreationReq{}
		err  error
	)

	if err = c.BodyParser(body); err != nil {
		response(c, http.StatusBadRequest, JSON{"message": "Error, parsing the body"})
		return
	}

	if err = h.validate.Struct(body); err != nil {
		response(c, http.StatusBadRequest, JSON{"message": "Error, body does not attend all requirements, ERROR:" + err.Error()})
		return
	}

	if acc, err = h.s.Create(body.Document); err != nil {
		if err == service.DocumentConflictError {
			response(c, http.StatusConflict, JSON{"message": service.DocumentConflictError.Error()})
			return
		}
		response(c, http.StatusInternalServerError, JSON{"message": err.Error()})
		return
	}

	response(c, http.StatusCreated, acc)
}

// FindById godoc
// @Summary Find an account
// @Description get account by its account id
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param accountID path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} handler.JSON
// @Failure 404 {object} handler.JSON
// @Failure 500 {object} handler.JSON
// @Router /accounts/{accountID} [get]
func (h *AccountsHandler) FindById(c *fiber.Ctx) {
	var (
		idSTR     = c.Params("accountID")
		account   *model.Account
		accountID int
		err       error
	)

	if accountID, err = strconv.Atoi(idSTR); err != nil {
		response(c, http.StatusBadRequest, JSON{"message": "AccountID is in the wrong format"})
		return
	}

	if account, err = h.s.FindById(accountID); err != nil {
		response(c, http.StatusNotFound, JSON{"message": "Account not found"})
		return
	}

	response(c, http.StatusOK, account)
}

// FindAll godoc
// @Summary List all accounts on repository
// @Description find accounts
// @Tags accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Account
// @Failure 500 {object} handler.JSON
// @Router /accounts [get]
func (h *AccountsHandler) FindAll(c *fiber.Ctx) {
	var (
		accounts []model.Account
		err      error
	)

	if accounts, err = h.s.FindAll(); err != nil {
		response(c, http.StatusInternalServerError, JSON{"message": err.Error()})
		return
	}

	response(c, http.StatusOK, accounts)
}
