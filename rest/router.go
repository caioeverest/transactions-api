package rest

import (
	swagger "github.com/arsmn/fiber-swagger"
	_ "github.com/caioeverest/transactions-api/docs"
	"github.com/caioeverest/transactions-api/rest/handler"
	"github.com/caioeverest/transactions-api/service"
)

// @title Transactions-api
// @version 1.0
// @description Experiment API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email caioeverest.b@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func (s *server) route(accountsService service.AccountsInterface, transactionsInterface service.TransactionsInterface) {
	accHandler := handler.NewAccountsHandler(accountsService, validate)
	trxHandler := handler.NewTransactionsHandler(transactionsInterface, validate)

	s.Use("/docs", swagger.Handler)
	s.Get("/health", handler.Health)

	accounts := s.Group("/accounts")
	{
		accounts.Post("/", accHandler.Create)
		accounts.Get("/", accHandler.FindAll)
		accounts.Get("/:accountID", accHandler.FindById)
	}

	transactions := s.Group("/transactions")
	{
		transactions.Post("/", trxHandler.Create)
		transactions.Get("/", trxHandler.FindAll)
	}
}
