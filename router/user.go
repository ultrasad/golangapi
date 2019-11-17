package router

import (
	"golangapi/handler"
	"golangapi/models"

	gormdb "golangapi/db/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//InitialRouteUser is init user route
func InitialRouteUser(e *echo.Echo) {
	//user := handler.UserHandler{}

	//db := gormdb.ConnectMySQL()
	//h := handler.NewHandler(models.NewUserModel(db))
	//user := handler.NewHandler(models.NewUserModel(db))

	user := handler.NewUserHandler(models.NewUserModel(gormdb.DBManager()))

	e.GET("/user/:id", user.GetUserByID)

	e.POST("/user", user.CreateUser)

	//user test with manual db connect
	e.GET("/allusers", user.GetAllUser)

	//2019-11-16, new db workshop

	// Restricted group
	r := e.Group("/api")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/customers", user.GetAllCustomer)

	//e.GET("/customers", user.GetAllCustomer)

}
