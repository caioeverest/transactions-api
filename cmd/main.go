package main

import (
	"os"
	"os/signal"

	"github.com/caioeverest/transactions-api/config"
	"github.com/caioeverest/transactions-api/logger"
	"github.com/caioeverest/transactions-api/repository"
	"github.com/caioeverest/transactions-api/rest"
	"github.com/caioeverest/transactions-api/service"
)

func main() {
	config.Start()
	logger.Start()

	operationsRepo,
		accountsRepo,
		transactionsRepo := repository.Start()

	accountsService,
		transactionsService := service.Start(operationsRepo, accountsRepo, transactionsRepo)

	rest.Start(accountsService, transactionsService)

	gracefulShutdown(
		repository.Shutdown,
		//rest.Shutdown,
	)
}

func gracefulShutdown(functions ...func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	for _, function := range functions {
		function()
	}

	os.Exit(0)
}
