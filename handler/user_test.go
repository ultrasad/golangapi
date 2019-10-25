package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/mock"

	//"golangapi/controllers"
	"golangapi/models"
)

/*
var (

	//now = time.Now().UTC()
	layout = "2006-01-02T15:04:05.000Z"
	str    = "2014-11-12T11:45:26.371Z"
	t, err = time.Parse(time.RFC3339, str)

	mockDB = map[string]*models.User{
		//"Hanajung@labstack.com": &models.User{Prefix: "Mr", Name: "Bundit", Email: "ultrasad@gmail.com", CreateDate: "2019-10-21"},
		"test": &models.User{ID: 0, Prefix: "Mr", Name: "Bundit", Email: "ultrasad@gmail.com", CreateDate: "2019-10-21", Timestamp: t},
	}

	//userJSON = `{"id":0, "prefix": "Mr", "name":"Bundit", "email":"ultrasad@gmail.com", "create_date":"2019-10-21", "timestamp":" + now + "}`
	userJSON = fmt.Sprintf(`{"id":0,"prefix":"Mr","name":"Bundit","email":"ultrasad@gmail.com","create_date":"2019-10-21","timestamp":"%s"}%s`, t.Format(layout), "\n")
)
*/

var (
	layout = "2006-01-02T15:04:05.000Z"
	str    = "2014-11-12T11:45:26.371Z"
	t, err = time.Parse(time.RFC3339, str)
)

type (
	UsersModelStub struct{}
)

func (u *UsersModelStub) GetUserByID(id string) models.User {
	return models.User{
		ID:         1,
		Prefix:     "Mr",
		Name:       "Hanajung",
		Email:      "kissing-bear@hotmail.com",
		CreateDate: "2019-10-24",
		Timestamp:  t,
	}
}

func (u *UsersModelStub) GetAllUser() []models.User {
	users := []models.User{}

	users = append(users, models.User{
		ID:         1,
		Prefix:     "Mr",
		Name:       "Hanajung",
		Email:      "kissing-bear@hotmail.com",
		CreateDate: "2019-10-24",
		Timestamp:  t,
	})

	fmt.Println("users => ", users)

	return users
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestGetUser(t *testing.T) {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//h := &userHandler{mockDB}
	//u := &UsersModelStub{}
	//h := &userHandler{u}

	//var u *UsersModelStub
	//u = NewSalutation()

	//u := UsersModelStub{1, "xx", "test", "mm"}
	//var h = &UserHandler{}

	u := &UsersModelStub{}
	h := NewUserHandler(u)

	var userJSON = fmt.Sprintf(`{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")

	//fmt.Println("String => ", rec.Body.String())

	// Assertions
	if assert.NoError(t, h.GetUserByID(c)) {
		//if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

/* func TestGetIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/allusers")

	u := &UsersModelStub{}
	h := NewHandler(u)

	var userJSON = `{"users":[{"id":100,"name":"foo"}]}`

	if assert.NoError(t, h.GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
} */

func TestGetAllUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/allusers")

	u := &UsersModelStub{}
	h := NewUserHandler(u)

	//var userJSON = `{"users":[{"id":100,"name":"foo"}]}`
	var userJSON = fmt.Sprintf(`[{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}]%s`, "\n")

	//fmt.Println("Say hi")
	//fmt.Printf("Body => %s", rec.Body.String())
	//t.Log("=> ", rec.Body.String())

	if assert.NoError(t, h.GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

/*
func TestHTTPGetUsers(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		//handler := http.HandlerFunc(handler.UserHandler)
		//user := UserHandler{}

		handler := http.HandlerFunc()
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}
*/
