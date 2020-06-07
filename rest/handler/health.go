package handler

import (
	"net/http"

	"github.com/caioeverest/transactions-api/config"
	"github.com/caioeverest/transactions-api/repository"
	"github.com/gofiber/fiber"
)

type HeartbeatResponse struct {
	Greetings string `json:"greetings"`
	Stage     string `json:"stage"`
	Database  string `json:"database"`
}

// Health godoc
// @Summary Application health check
// @Description return the status of the application and connectivity with the database
// @Tags general
// @Accept  json
// @Produce  json
// @Success 200 {array} handler.HeartbeatResponse
// @Failure 500 {object} handler.JSON
// @Router /health [get]
func Health(c *fiber.Ctx) {
	var dbStatus = "OFF"

	if ok := repository.Health(); ok {
		dbStatus = "OK"
	}

	response(c, http.StatusOK, HeartbeatResponse{
		Greetings: config.Get().HTTP.Greetings,
		Database:  dbStatus,
		Stage:     config.Get().ENV,
	})
}
