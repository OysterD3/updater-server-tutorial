package router

import (
	"context"

	"github.com/OysterD3/updater-server-tutorial/bootstrap"
	"github.com/labstack/echo/v4"
)

// New :
func New(e *echo.Group) *echo.Group {
	bs := bootstrap.New(context.Background())
	InitRoutes(e, bs)
	return e
}
