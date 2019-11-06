package router

import (
	"golangapi/db/mongo"
	"golangapi/handler"

	"github.com/labstack/echo/v4"
)

//InitialRouteDoc is init doc route
func InitialRouteDoc(e *echo.Echo) {

	//new mongo connect
	store := mongo.DataStoreNew()

	//new test document
	doc := handler.NewDocumentHandler(store)
	//doc := handler.DocHandler{}
	e.GET("/docs", doc.Test)
	e.GET("/find_docs", doc.Find)

	//FindWithMgo
	e.GET("/find_doc_mgo/:docID", doc.FindWithMgo)

	//e.PUT("/todos/:id", todo.UpdateTodo) //update, done

	/* e.PUT("/check/:version", func(c echo.Context) error {
		version := c.Param("version")
		//return c.String(http.StatusOK, version)
		return c.JSON(http.StatusOK, version)
	}) */

	/* e.PUT("/todos/:idx", func(c echo.Context) error {
		idx := c.Param("idx")
		fmt.Println("idx => ", idx)
		return c.String(http.StatusOK, idx)
	}) */
	//e.DELETE("/todos/:id", todo.DeleteTodo)
}
