package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/mock"
	//"golangapi/controllers"
	//"golangapi/models"
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
	//strString   = "2019-11-04T11:45:26.371Z"
	tt, _ = time.Parse(time.RFC3339, str)
	//ttString, _ = time.Parse("2006-01-02", strString)

	/* mockDB = map[string]*models.User{
		"jon@labstack.com": &models.User{"Jon Snow", "jon@labstack.com"},
	} */

	mockUserDB = models.User{
		ID:     1,
		Prefix: "Mr",
		Name:   "Hanajung",
		Email:  "kissing-bear@hotmail.com",
		//CreateDate: "2019-10-24",
		CreateDate: time.Now().Local(),
		Timestamp:  tt,
	}

	//userMockJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
)

type (
	UsersModelStub struct{}
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()

	//var userJSON = fmt.Sprintf(`{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")
	var userMockJSONResponse = fmt.Sprintf(`{"id":3,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")
	var userJSON = fmt.Sprintf(`{"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com", "create_date":"2019-10-24", "timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	//h := &handler{mockDB}
	u := &UsersModelStub{}
	h := NewUserHandler(u)

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userMockJSONResponse, rec.Body.String())
	}
}

func (u *UsersModelStub) CreateUserWithTransection(user *models.User) (*models.User, error) {
	return &models.User{
		ID:     3,
		Prefix: "Mr",
		Name:   "Hanajung",
		Email:  "kissing-bear@hotmail.com",
		//CreateDate: "2019-10-24",
		//CreateDate: time.Now().Local(),
		CreateDateString: "2019-10-24",
		Timestamp:        tt,
	}, nil
	//return mockUserDB
}

/* func (u *UsersModelStub) CreateUser(id string) models.User {
	return models.User{
		ID:         1,
		Prefix:     "Mr",
		Name:       "Hanajung",
		Email:      "kissing-bear@hotmail.com",
		CreateDate: "2019-10-24",
		Timestamp:  tt,
	}
	//return mockUserDB
} */

func (u *UsersModelStub) GetUserByID(id string) models.User {
	return models.User{
		ID:               1,
		Prefix:           "Mr",
		Name:             "Hanajung",
		Email:            "kissing-bear@hotmail.com",
		CreateDate:       tt,
		CreateDateString: "2019-11-04",
		//Timestamp:  time.Time,
		//Timestamp: models.CustomTime{t},
		Timestamp: tt,
	}
}

func (u *UsersModelStub) GetAllUser() []models.User {
	users := []models.User{}

	users = append(users, models.User{
		ID:     1,
		Prefix: "Mr",
		Name:   "Hanajung",
		Email:  "kissing-bear@hotmail.com",
		//CreateDate: "2019-10-24",
		//CreateDate: time.Now().Local(),
		CreateDate:       tt,
		CreateDateString: "2019-11-04",
		//Timestamp:  models.CustomTime{t},
		Timestamp: tt,
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

	//var userJSON = fmt.Sprintf(`{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2019-10-24","timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")
	var userJSON = fmt.Sprintf(`{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2014-11-12","timestamp":"2014-11-12T11:45:26.371Z"}%s`, "\n")

	//fmt.Println("String => ", rec.Body.String())

	// Assertions
	if assert.NoError(t, h.GetUserByID(c)) {
		//if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

/* func TestCreateUser(t *testing.T) {
	e := echo.New()
	user := NewUserHandler(models.NewUserModel(gormdb.DBManager()))

	testUser := &models.User{
		ID:         3,
		Prefix:     "Mr",
		Name:       "Hanajung",
		Email:      "kissing-bear@hotmail.com",
		CreateDate: "2019-10-24",
		Timestamp:  tt,
	}

	// Transform Star record into *strings.Reader suitable for use in HTTP POST forms.
	data := url.Values{
		"prefix":      {testUser.Prefix},
		"name":        {testUser.Name},
		"email":       {testUser.Email},
		"create_date": {testUser.CreateDate},
	}

	form := strings.NewReader(data.Encode())

	// Set up a new request.
	req, err := http.NewRequest("POST", "/stars", form)
	if err != nil {
		t.Fatal(err)
	}

	// Our API expects a form body, so set the content-type header to make sure it's treated as one.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//rr := httptest.NewRecorder()
	//http.HandlerFunc(user.CreateUser).ServeHTTP(rr, req)

	//req := httptest.NewRequest(echo.POST, "/", form)
	//req.Header.Add("Content-Type", "application/json")
	//rec := httptest.NewRecorder()
	//c := e.NewContext(req, rec)
	//c.SetPath("/users")

	//if assert.NoError(t, user.CreateUser(c)) {
	//	assert.Equal(t, http.StatusOK, rec.Code)
	//	assert.Equal(t, userJSON, rec.Body.String())
	//}
} */

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
	var userJSON = fmt.Sprintf(`[{"id":1,"prefix":"Mr","name":"Hanajung","email":"kissing-bear@hotmail.com","create_date":"2014-11-12","timestamp":"2014-11-12T11:45:26.371Z"}]%s`, "\n")

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
