package _handler

import (
	"fmt"
	"golangapi/models"
	"net/http"
	"strconv"

	//"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo/v4"
)

//TodoHandler is Todo
type TodoHandler struct {
	Todo *models.Todo
}

// List todo
func (h *TodoHandler) List(c echo.Context) (err error) {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	result, err := models.FindAllTodos(page, limit)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Create todo
func (h *TodoHandler) Create(c echo.Context) (err error) {
	id := bson.NewObjectId()
	var t models.Todo
	if err := c.Bind(&t); err != nil {
		return err
	}

	t.ID = id
	t.Done = false

	result, err := models.CreateTodo(&t)
	return c.JSON(http.StatusOK, result)
}

// View todo
func (h *TodoHandler) View(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))
	//id := c.Param("id")
	result, err := models.FindTodoByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Done todo
func (h *TodoHandler) Done(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))
	//id := c.Param("id")
	var t models.Todo

	t, err = models.FindTodoByID(id)
	if err != nil {
		return err
	}

	fmt.Println("before bind data controller => id, data => ", id, &t)

	if err := c.Bind(&t); err != nil {
		return err
	}

	t.Done = true
	fmt.Printf("new todo update done: %+v\n", &t)

	result, err := models.UpdateTodo(id, &t)
	fmt.Println("after bind data controller => id, data => ", id, &t)

	//return c.JSON(http.StatusOK, map[string]string{"result": "success"})
	return c.JSON(http.StatusOK, result)
}

//Delete todo
func (h *TodoHandler) Delete(c echo.Context) (err error) {
	id := bson.ObjectIdHex(c.Param("id"))
	//id := c.Param("id")
	err = models.DeleteTodo(id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
	return nil
}
