package handler

import (
	"net/http"

	"github.com/caioeverest/transactions-api/logger"
	"github.com/gofiber/fiber"
)

type JSON map[string]interface{}

func response(c *fiber.Ctx, status int, data interface{}) {
	if err := c.Status(status).JSON(data); err != nil {
		logger.Get().Errorf("Error parsing return object - ERROR [%+v]", err)
		c.Status(http.StatusInternalServerError).Send(err.Error())
	}
}
