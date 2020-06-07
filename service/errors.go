package service

import "github.com/caioeverest/transactions-api/util"

//Accounts Errors
const (
	DocumentConflictError = util.Error("this document are already associated with an account")
)

//Transactions Errors
const (
	AmountCantBeNegativeError = util.Error("amount must always be a positive value")
	AccountNotFoundError      = util.Error("can't find this account id on repository")
	OperationNotFoundError    = util.Error("can't find this operation id on repository")
)
