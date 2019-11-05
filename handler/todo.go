package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"golangapi/models"

	"github.com/labstack/echo/v4"
)

type (
	//TodoHandler is Todo Controller with model
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

	//fmt.Println("call handler time => ", time.Now().Local())

	id := c.Param("todoID")
	result, err := h.TodoModel.GetTodo(id)
	if err != nil {
		fmt.Println("GetTodo Error: ", err)
		return c.JSON(http.StatusBadRequest, err)
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
	todoID := c.Param("todoID")

	fmt.Printf("Update topic todoID: %s \n", todoID)
	var todo models.Todo

	todo, err = h.TodoModel.GetTodo(todoID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	//todo.ID = idx.Hex()
	//fmt.Printf("before bind data controller data %v\n ", &todo)

	if err := c.Bind(&todo); err != nil {
		fmt.Println("BindTodo Error, ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	//fmt.Println("\n todo type:", reflect.TypeOf(todo))
	//fmt.Printf("after bind data controller data %v\n ", &todo)

	//todo.ID = idx
	//todo.Fake = true
	//todo.Topic = topic
	//fmt.Printf("new todo update done: %+v\n", &todo)

	result, err := h.TodoModel.UpdateTodo(todoID, &todo)
	//return c.JSON(http.StatusOK, map[string]string{"result": "success"})
	return c.JSON(http.StatusOK, result)
}

//DeleteTodo is delete todo by id
func (h *TodoHandler) DeleteTodo(c echo.Context) (err error) {
	id := c.Param("todoID")

	zapLog := ZapManager()
	zapLog.Info("delete topic todoID: " + id)

	deleteID, err := h.TodoModel.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"result": "success",
		"id":     deleteID,
	})
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
