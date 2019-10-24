package router

import (
	"golangapi/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//InitialRouteTodo is init todo route
func InitialRouteTodo(e *echo.Echo) {
	todo := handler.TodoHandler{}

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/todos", todo.List)

	e.GET("/todos", todo.List)
	e.POST("/todos", todo.Create)
	e.GET("/todos/:id", todo.View)
	e.PUT("/todos/:id", todo.Done)
	e.DELETE("/todos/:id", todo.Delete)
}
