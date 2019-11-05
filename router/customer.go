package router

import (
	"golangapi/handler"

	"github.com/labstack/echo/v4"
)

//InitialRouteCustomer is init todo route
func InitialRouteCustomer(e *echo.Echo) {
	customer := handler.NewCustomerHandler()
	e.GET("/essearch", customer.Search)
}
