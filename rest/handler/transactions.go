package handler

import (
	"net/http"

	"github.com/caioeverest/transactions-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

type TransactionsHandler struct {
	s        service.TransactionsInterface
	validate *validator.Validate
}

type TransactionCreationReq struct {
	AccountID       int     `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

func NewTransactionsHandler(s service.TransactionsInterface, v *validator.Validate) *TransactionsHandler {
	return &TransactionsHandler{s, v}
}

// Create godoc
// @Summary Create a new transaction
// @Description create a new transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param account body handler.TransactionCreationReq true "Create a transaction"
// @Success 200 {object} service.TransactionDTO
// @Failure 400 {object} handler.JSON
// @Failure 500 {object} handler.JSON
// @Router /transactions [post]
func (h *TransactionsHandler) Create(c *fiber.Ctx) {
	var (
		trx  *service.TransactionDTO
		body = &TransactionCreationReq{}
		err  error
	)

	if err = c.BodyParser(body); err != nil {
		response(c, http.StatusBadRequest, JSON{"message": "body does not have all the expected parameters", "error": err})
		return
	}

	if err = h.validate.Struct(body); err != nil {
		response(c, http.StatusBadRequest, JSON{"message": "body does not have all the expected parameters", "error": err})
		return
	}

	if trx, err = h.s.Create(body.AccountID, body.OperationTypeId, body.Amount); err != nil {
		switch err {
		case service.AmountCantBeNegativeError:
			fallthrough
		case service.OperationNotFoundError:
			response(c, http.StatusBadRequest, JSON{"message": err.Error()})
		case service.AccountNotFoundError:
			response(c, http.StatusNotFound, JSON{"message": err.Error()})
		default:
			response(c, http.StatusInternalServerError, JSON{"message": err.Error()})
		}
		return
	}

	response(c, http.StatusCreated, trx)
}

// FindAll godoc
// @Summary List all transactions on repository
// @Description list transactions
// @Tags transactions
// @Accept  json
// @Produce  json
// @Success 200 {array} service.TransactionDTO
// @Failure 500 {object} handler.JSON
// @Router /transactions [get]
func (h *TransactionsHandler) FindAll(c *fiber.Ctx) {
	var (
		transactions []service.TransactionDTO
		err          error
	)

	if transactions, err = h.s.FindAll(); err != nil {
		response(c, http.StatusInternalServerError, JSON{"message": err.Error()})
		return
	}

	response(c, http.StatusOK, transactions)
}
