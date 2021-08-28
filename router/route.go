package router

import (
	"github.com/OysterD3/updater-server-tutorial/bootstrap"
	"github.com/OysterD3/updater-server-tutorial/handler"
	"github.com/labstack/echo/v4"
)

// InitRoutes :
func InitRoutes(e *echo.Group, bs *bootstrap.Bootstrap) (api *echo.Group) {
	h := handler.New(bs)
	api = e.Group("")

	api.GET("/", h.HealthCheck)
	api.GET("/update/:platform/:version", h.GetReleasesByPlatformAndVersion)
	api.GET("/download/:platform/:version", h.Download)

	return
}
