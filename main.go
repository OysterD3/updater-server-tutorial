package main

import (
	"github.com/OysterD3/updater-server-tutorial/env"
	"github.com/OysterD3/updater-server-tutorial/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	api := e.Group("")
	router.New(api)

	e.Logger.Fatal(e.Start(":" + env.Config.Port))
}
