package handler

import (
	"fmt"
	"golangapi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	//TodoHandler is Todo
	TodoHandler struct {
		TodoModel models.TodoStore
	}
)

//NewTodoHandler is new todo
func NewTodoHandler(u models.TodoStore) *TodoHandler {
	return &TodoHandler{u}
}

//GetTodo reponse todo by id, (//FindTodo)
func (h *TodoHandler) GetTodo(c echo.Context) error {
	id := c.Param("id")
	result, err := h.TodoModel.GetTodo(id)
	if err != nil {
		fmt.Println("GetTodo Error: ", err)
	}

	return c.JSON(http.StatusOK, result)
}

//CreateTodo create new todo
func (h *TodoHandler) CreateTodo(c echo.Context) error {
	todo := models.Todo{}

	if err := c.Bind(&todo); err != nil {
		fmt.Println("BindTodo Error, ", err)
		//return c.NoContent(http.StatusBadRequest)
		//return json response
	}

	result, err := h.TodoModel.CreateTodo(&todo)
	if err != nil {
		fmt.Println("CreateTodo Error: ", err)
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateTodo is update todo by id
func (h *TodoHandler) UpdateTodo(c echo.Context) (err error) {
	//id := bson.ObjectIdHex(c.Param("id"))
	id := c.Param("id")
	//todo := models.Todo{}
	var todo models.Todo
	//todo := new(models.Todo)

	todo, err = h.TodoModel.GetTodo(id)
	if err != nil {
		fmt.Println("GetTodo Error: ", err)
		//return c.NoContent(http.StatusNotFound)
		//return json response
	}

	fmt.Println("before bind data controller => id, data => ", id, &todo)

	/* if err := c.Bind(todo); err != nil {
		fmt.Println("BindTodo Error, ", err)
		//return c.NoContent(http.StatusBadRequest)
		//return json response
	} */

	if err := c.Bind(&todo); err != nil {
		return err
	}

	todo.Done = true
	fmt.Printf("new todo update done: %+v\n", &todo)

	result, err := h.TodoModel.UpdateTodo(id, &todo)
	fmt.Println("after bind data controller => id, data => ", id, &todo)

	//return c.JSON(http.StatusOK, map[string]string{"result": "success"})
	return c.JSON(http.StatusOK, result)
}

//GetAllTodo response all todo with limit, perpage
func (h *TodoHandler) GetAllTodo(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	result, err := h.TodoModel.GetAllTodo(int64(page), int64(limit))
	if err != nil {
		fmt.Println("GetAllTodo Error: ", err)
	}

	return c.JSON(http.StatusOK, result)
}
