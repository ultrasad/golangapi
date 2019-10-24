package router

import (
	"golangapi/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//InitialRouteAuth is init auth route
func InitialRouteAuth(e *echo.Echo) {
	auth := handler.AuthHandler{}

	// Login route
	e.POST("/login", auth.Login)

	// Unauthenticated route
	e.GET("/", auth.Accessible)

	// Restricted group
	r := e.Group("/private")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/restricted", auth.Restricted)
}
