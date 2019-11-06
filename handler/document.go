package handler

import (
	"context"
	"net/http"

	mongoClient "golangapi/db/mongo"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	//Document is doc test
	Document struct {
		Topic string `json:"topic" bson:"topic"`
		Done  bool   `json:"done" bson:"done"`
	}

	//TestRepository test
	/* TestRepository interface {
		Find(ctx echo.Context, filters interface{}) ([]Document, error)
	} */

	//DocHandler is new doc
	DocHandler struct {
		store *mongoClient.DBNewDataStore
		//store *mongoClient.DBNewDataStoreClient
		//client *mongo.Client
	}

	/* //TodoXHandler is Todo Controller with model
	TodoXHandler struct {
		TodoModel models.TodoStore
	} */
)

//Test ...
func (h *DocHandler) Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "Test Doc")
}

/* //NewTodoXHandler is new todo
func NewTodoXHandler(u models.TodoStore) *TodoHandler {
	return &TodoHandler{u}
}*/

//NewDocumentHandler is new doc
func NewDocumentHandler(u *mongoClient.DBNewDataStore) *DocHandler {
	return &DocHandler{u}
}

//Find ...
func (h *DocHandler) Find(c echo.Context) error {

	//cur, err := r.store.GetCollection("some_collection_name").Find(ctx, filters)
	//var filters interface{}

	//h.store = mongoClient.DataStoreNew()

	filters := bson.M{}
	findOptions := options.Find()
	findOptions.SetLimit(10)
	findOptions.SetSkip(0)
	findOptions.SetSort(bson.M{"_id": -1})

	ctx := context.TODO()
	//cur, err := h.store.Session.Database("document").Collection("todo").Find(ctx, filters)
	/* curSess := h.store.Session

	err := curSess.Ping(ctx, nil)

	if err != nil {
		fmt.Println("ping mongodb err,", err)
		//log.Fatal(err)
	} */

	//fmt.Println("mongo store => ", h.store)

	//newclient := mongoClient.ClientNewManager()
	//newclient := mongoClient.ClientManager()

	cur, err := h.store.Session.Database("document").Collection("todo").Find(ctx, filters, findOptions)
	//cur, err := newclient.Database("document").Collection("todo").Find(ctx, filters, findOptions)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	//var result = make([]models.Document, 0)
	var result = make([]Document, 0)
	for cur.Next(ctx) {
		//var currDoc models.Document
		var currDoc Document
		err := cur.Decode(&currDoc)
		if err != nil {
			//log here
			continue
		}
		result = append(result, currDoc)
	}
	//return err
	return c.JSON(http.StatusOK, result)
}
