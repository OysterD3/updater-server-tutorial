package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
