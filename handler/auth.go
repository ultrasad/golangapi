package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type (
	//AuthHandler is auth struct
	AuthHandler struct{}
)

//Accessible ...
func (h *AuthHandler) Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

//Restricted ...
func (h *AuthHandler) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	//return c.String(http.StatusOK, "Welcome "+name+"!")
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Welcome " + name,
	})
}

//Login ...
func (h *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	//email := c.FormValue("email")
	password := c.FormValue("password")

	//fmt.Println("login => ", username, password)

	// in our case, the only "valid user and password" is
	// user: rickety_cricket@example.com pw: shhh!
	// really, this would be connected to any database and
	// retrieving the user and validating the password
	//if email != "rickety_cricket@example.com" || password != "shhh!" {
	if username != "hanajung" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	//fmt.Println("token => ", token)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	// add any key value fields to the token
	//claims["email"] = "rickety_cricket@example.com"
	claims["name"] = "Hanajung"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// return the token for the consumer to grab and save
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
