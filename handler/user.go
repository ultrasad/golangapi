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

/* var (
	layout = "2006-01-02T15:04:05.000Z"
	str    = "2014-11-12T11:45:26.371Z"
	t, err = time.Parse(time.RFC3339, str)
) */

type (
	//UserHandler is new user handler
	UserHandler struct {
		//DB *gorm.DB
		//user *models.User
		UserModel models.UserModelImpl
	}

	//User ...
	/* User struct {
		//BaseModel
		ID         uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		Prefix     string    `json:"prefix"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		CreateDate string    `json:"create_date"`
		Timestamp  time.Time `json:"timestamp" gorm:"column:timestamp" sql:"DEFAULT:current_timestamp"`
	} */
)

//NewHandler ...
func NewHandler(u models.UserModelImpl) *UserHandler {
	return &UserHandler{u}
}

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

//GetUserByID is get user by id
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	result := h.UserModel.GetUserByID(id)
	//fmt.Println("json response => ", result)

	/* user := User{
		ID:         1,
		Prefix:     "Mr",
		Name:       "Hanajung",
		Email:      "kissing-bear@hotmail.com",
		CreateDate: "2019-10-24",
		Timestamp:  t,
	}

	js, _ := json.Marshal(user)
	return c.JSONBlob(http.StatusOK, js) */

	return c.JSON(http.StatusOK, result)
}

//GetAllUser is get all user
func (h *UserHandler) GetAllUser(c echo.Context) error {

	//result := models.GetUsers()
	result := h.UserModel.GetAllUser()

	for i, ar := range result {
		//createDate, _ := time.Parse("2006-01-02T00:00:00Z", ar.CreateDate)
		createDate, _ := time.Parse("2006-01-02", ar.CreateDate)
		/* if error != nil {
			fmt.Println("error => ", error)
		} */

		fmt.Println("reponse date => ", createDate.Format("2006-01-02"))
		result[i].CreateDate = createDate.Format("2006-01-02")
	}

	return c.JSON(http.StatusOK, result)
	//return c.JSON(200, `[{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}]`)
}

/* func (h *UserHandler) GetAllUser(c echo.Context) error {

	//result := models.GetUsers()
	result := h.UserModel.GetAllUser()

	for i, ar := range result.Users {
		createDate, _ := time.Parse("2006-01-02T00:00:00Z", ar.CreateDate)
		result.Users[i].CreateDate = createDate.Format("2006-01-02")
	}

	return c.JSON(http.StatusOK, result)
} */

//GetUserDefault is get all user list
func (h *UserHandler) GetUserDefault(c echo.Context) error {
	fmt.Println("call get all user")
	result := models.GetUserDefault()
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
