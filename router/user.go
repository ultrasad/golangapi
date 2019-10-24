package router

import (
	"golangapi/handler"
	"golangapi/models"

	gormdb "golangapi/db/gorm"

	"github.com/labstack/echo/v4"
)

//InitialRouteUser is init user route
func InitialRouteUser(e *echo.Echo) {
	//user := handler.UserHandler{}

	//db := gormdb.ConnectMySQL()
	//h := handler.NewHandler(models.NewUserModel(db))
	//user := handler.NewHandler(models.NewUserModel(db))

	user := handler.NewHandler(models.NewUserModel(gormdb.DBManager()))

	e.GET("/users/:id", user.GetUserByID)

	//user test with manual db connect
	e.GET("/allusers", user.GetAllUser)
}
