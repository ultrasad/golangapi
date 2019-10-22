package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//ServerHeader is serve header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func initEcho() *echo.Echo {
	e := echo.New()
	e.Use(ServerHeader)
	e.GET("/", rootURL)
	return e
}

func rootURL(c echo.Context) error {
	return c.HTML(http.StatusOK, "body output")
}
