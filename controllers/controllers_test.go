package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	//"golangapi/controllers"
	"golangapi/models"
)

var (
	/*
		mockDB = map[string]*User{
			"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
		}

		userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
	*/

	/*mockDB = map[string]*models.User{
		"Hanajung@labstack.com": &models.User{Prefix: "Mr", Name: "Bundit", Email: "ultrasad@gmail.com", CreateDate: "2019-10-21"},
	}*/

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

func TestGetUser(t *testing.T) {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//c.SetPath("/users/:email")
	//c.SetParamNames("email")
	//c.SetParamValues("jon@labstack.com")
	//h := &models.User{mockDB}
	//m := mockDB
	//h := &controllers

	//fmt.Println("t => ", t)

	h := &userModel{mockDB}

	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("test")

	// Assertions
	if assert.NoError(t, h.GetUserMock(c)) {
		//if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

/*
func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.createUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")
	h := &handler{mockDB}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
*/
