package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

type (
	UsersModelStub struct {
		//mock.Mock
	}
)

func (u *UsersModelStub) GetUser(id string) models.User {
	return models.User{
		ID:   1,
		Name: "foo",
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
	c.SetParamValues("test")

	//h := &userHandler{mockDB}
	//u := &UsersModelStub{}
	//h := &userHandler{u}

	h := &UserHandler{}

	var userJSON = `{"id":1,"name":"foo"}`

	// Assertions
	if assert.NoError(t, h.GetUser(c)) {
		//if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
