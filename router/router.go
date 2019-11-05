package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//InitialRoute is init route
func InitialRoute(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	//Init auth
	InitialRouteAuth(e)

	//Init todo
	InitialRouteTodo(e)

	//Init user
	InitialRouteUser(e)

	//Init user
	InitialRouteCustomer(e)

	fmt.Println("InitialRoute...")

	//db with gorm

	//GoRoutine
	//e.GET("/hello", controllers.CallHelloRoutine)

	//Elastics Route
	//e.GET("/esversion", controllers.ESVersion)

	//Elastics Search
	//e.GET("/essearch", controllers.Search)

}
