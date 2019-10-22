package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golangapi/models"

	//"github.com/labstack/echo"

	"github.com/labstack/echo/v4"
)

type (
	//UserHandler is new user handler
	UserHandler struct {
		//DB *gorm.DB
		//user *models.User
		UserModel models.UserModelImpl
	}
)

//CreateUser is create new user
func (h *UserHandler) CreateUser(c echo.Context) (err error) {

	jsonMap := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}

	//jsonMap has the JSON Payload decoded into a map
	prefix := jsonMap["prefix"]
	name := jsonMap["name"]
	email := jsonMap["email"]
	inputdate := jsonMap["create_date"]

	inputtamp := jsonMap["timestamp"]

	//"2006-01-02 15:04:05" is standard format datetime golang
	timestamp, _ := time.Parse("2006-01-02 15:04:05", inputtamp.(string))

	user := &models.User{Name: name.(string), Email: email.(string), Prefix: prefix.(string), Timestamp: timestamp, CreateDate: inputdate.(string)}

	//user.Timestamp.Format("2006-01-02 15:04:05")
	err = models.CreateUser(user)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, user)
}

//GetUsers is get user
func (h *UserHandler) GetUsers(c echo.Context) error {

	result := models.GetUsers()

	for i, ar := range result.Users {
		createDate, _ := time.Parse("2006-01-02T00:00:00Z", ar.CreateDate)
		result.Users[i].CreateDate = createDate.Format("2006-01-02")
	}

	return c.JSON(http.StatusOK, result)
}

//GetAllUser is get all user list
func (h *UserHandler) GetAllUser(c echo.Context) error {
	fmt.Println("call get all user")
	result := models.GetUserDefault()
	return c.JSON(http.StatusOK, result)
}

//GetUser is get user by id
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	//result := models.GetUser(id)
	result := h.UserModel.GetUser(id)
	return c.JSON(http.StatusOK, result)
}

/*
func (h *userHandler) GetUserMock(c echo.Context) error {
	id := c.Param("id")
	user := h.user[id]

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	//fmt.Println("user => ", user)
	return c.JSON(http.StatusOK, user)
}
*/
