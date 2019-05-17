package controllers

import (
	"fmt"
	"golangapi/models"
	"net/http"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/labstack/echo"
)

// List todo
func List(c echo.Context) (err error) {
	result, err := models.FindAllTodos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Create todo
func Create(c echo.Context) (err error) {
	//id := bson.NewObjectId()
	var t models.Todo
	if err := c.Bind(&t); err != nil {
		return err
	}

	//t.ID = id
	t.Done = false

	//fmt.Println("todo => : ", &t)
	//return err

	result, err := models.CreateTodo(&t)
	return c.JSON(http.StatusOK, result)
}


// View todo
func View(c echo.Context) (err error) {
	//id := bson.ObjectIdHex(c.Param("id"))
	id := c.Param("id")
	result, err := models.FindTodoByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Done todo
func Done(c echo.Context) (err error) {
	//id := bson.ObjectIdHex(c.Param("id"))
	id := c.Param("id")
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

/*
// Update todo like done***
func Update(c echo.Context) (err error) {
	return err
}
*/

//Delete todo
func Delete(c echo.Context) (err error) {
	//id := bson.ObjectIdHex(c.Param("id"))
	id := c.Param("id")
	err = models.DeleteTodo(id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, echo.Map{
		"result": "success",
	})
	return nil
}