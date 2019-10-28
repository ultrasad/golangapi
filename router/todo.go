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
	e.POST("/todos", todo.CreateTodo)
	e.GET("/todos/:todoID", todo.GetTodo)
	e.PUT("/todos/:todoID", todo.UpdateTodo) //update, done

	//e.PUT("/todos/:id", todo.UpdateTodo) //update, done

	/* e.PUT("/check/:version", func(c echo.Context) error {
		version := c.Param("version")
		//return c.String(http.StatusOK, version)
		return c.JSON(http.StatusOK, version)
	}) */

	/* e.PUT("/todos/:idx", func(c echo.Context) error {
		idx := c.Param("idx")
		fmt.Println("idx => ", idx)
		return c.String(http.StatusOK, idx)
	}) */
	//e.DELETE("/todos/:id", todo.DeleteTodo)
}
