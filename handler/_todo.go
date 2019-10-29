package _handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"golangapi/models"

	"github.com/labstack/echo/v4"
)

type (
	//TodoHandler is Todo
	TodoHandler struct {
		TodoModel models.TodoStore
	}

	//CustomBinder ...
	CustomBinder models.Todo
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

// Bind ...
/* func (td *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	// You may use default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		fmt.Println("err bind => ", err)
		return
	}

	// Define your custom implementation
	fmt.Println("return bind => ", err)
	return
} */

// UpdateTodo is update todo by id
func (h *TodoHandler) UpdateTodo(c echo.Context) (err error) {
	todoID := c.Param("id")
	//other := c.Get("other")
	//topic := c.FormValue("topic")
	//topic := c.Param("topic")

	//zapLog := ZapManager()
	//zapLog.Info("Update topic todoID: " + todoID)

	fmt.Printf("Update topic todoID: %s \n", todoID)

	/* req := c.Request().Body

	names := c.ParamNames()
	values := c.ParamValues()
	params := map[string][]string{}
	for i, name := range names {
		params[name] = []string{values[i]}
	}

	fmt.Println("params => ", params, ", request => ", req) */

	//idx, _ := primitive.ObjectIDFromHex(todoID)
	//fmt.Printf("old todos => %s %v", err, id)
	//todo := models.Todo{}
	//var todo models.Todo
	//todo := new(models.Todo)

	//fmt.Println("\n todo type:", reflect.TypeOf(todo))
	//fmt.Println("\n idx type:", reflect.TypeOf(idx))

	//todo := make(map[string]interface{})
	//var todo map[string]interface{}
	var todo models.Todo

	/* todo, err = h.TodoModel.GetTodo(todoID)
	if err != nil {
		//fmt.Println("GetTodo Error: ", err)
		return c.JSON(http.StatusBadRequest, err)
		//return c.NoContent(http.StatusNotFound)
		//return json response
	} */

	if err := c.Bind(&todo); err != nil {
		return err
	}
	//return c.JSON(200, m)

	//todo.ID = idx.Hex()
	fmt.Printf("before bind data controller data %v\n ", &todo)

	/* if err = json.NewDecoder(req).Decode(&todo); err != nil {
		fmt.Println("NewDecoder Error, ", err)
	} */

	/* if err := c.Bind(&todo); err != nil {
		fmt.Println("BindTodo Error, ", err)
		//return c.NoContent(http.StatusBadRequest)
		//return json response
	} */

	fmt.Println("\n todo type:", reflect.TypeOf(todo))

	fmt.Printf("after bind data controller data %v\n ", &todo)

	//todo.ID = idx
	//todo.Fake = true
	//todo.Topic = topic
	//fmt.Printf("new todo update done: %+v\n", &todo)

	//result, err := h.TodoModel.UpdateTodo(todoID, &todo)
	//return c.JSON(http.StatusOK, map[string]string{"result": "success"})
	return c.JSON(http.StatusOK, err)
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
