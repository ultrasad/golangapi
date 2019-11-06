package router

import (
	"golangapi/handler"

	"github.com/labstack/echo/v4"
)

//InitialRouteCus is init todo route
func InitialRouteCus(e *echo.Echo) {
	customer := handler.NewCustomerHandler()
	e.GET("/essearch", customer.Search)
}
