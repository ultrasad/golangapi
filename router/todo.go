package router

import (
	"golangapi/db/mongo"
	"golangapi/handler"
	"golangapi/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//InitialRouteTodo is init todo route
func InitialRouteTodo(e *echo.Echo) {
	todo := handler.NewTodoHandler(models.NewTodoModel(mongo.ClientManager()))

	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	//r.GET("/todos", todo.List)

	e.GET("/todos", todo.GetAllTodo)
	e.GET("/todos/:id", todo.GetTodo)
	e.POST("/todos", todo.CreateTodo)
	e.PUT("/todos/:id", todo.UpdateTodo) //update, done
	//e.DELETE("/todos/:id", todo.DeleteTodo)
}
