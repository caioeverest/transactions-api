package rest

import (
	"github.com/caioeverest/transactions-api/config"
	"github.com/caioeverest/transactions-api/logger"
	"github.com/caioeverest/transactions-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

type server struct {
	*fiber.App
}

var (
	log      *logger.Logger
	api      *server
	validate *validator.Validate
)

// Start HTTP server
func Start(accountsService service.AccountsInterface, transactionsService service.TransactionsInterface) {
	conf := config.Get()
	log = logger.Get()
	validate = validator.New()
	api = &server{fiber.New()}
	api.Settings.DisableStartupMessage = true

	// Check if should set Fiber prefork config
	if !config.IsDevelopment() {
		api.Settings.Prefork = true
	}

	log.Info("Starting HTTP server...")
	api.route(accountsService, transactionsService)

	go func() {
		if err := api.Listen(conf.HTTP.Port); err != nil {
			log.Panicf("Error starting server - ERROR [%+v]", err)
		}
	}()
}

// Shutdown HTTP server
func Shutdown() {
	if err := api.Shutdown(); err != nil {
		log.Errorf("Error stopping server - ERROR [%+v]", err)
	}
	log.Info("Server shutting down")
}
