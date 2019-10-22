package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//InitialRoute is init route
func InitialRoute(e *echo.Echo) {

	//e := echo.New()
	todo := TodoHandler{}

	e.GET("/todos/:id", todo.View)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/restricted", restricted)
	r.GET("/todos", todo.List)

	e.GET("/todos", todo.List)
	e.POST("/todos", todo.Create)
	e.GET("/todos/:id", todo.View)
	e.PUT("/todos/:id", todo.Done)
	e.DELETE("/todos/:id", todo.Delete)

	fmt.Println("InitialRoute...")

	//db with gorm
	user := UserHandler{}
	e.GET("/users/:id", user.GetUser)

	//db local config
	//e.GET("/allusers", controllers.GetAllUser)

	//GoRoutine
	//e.GET("/hello", controllers.CallHelloRoutine)

	//Elastics Route
	//e.GET("/esversion", controllers.ESVersion)

	//Elastics Search
	//e.GET("/essearch", controllers.Search)

}
