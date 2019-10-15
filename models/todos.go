package models

import (
	mongo "golangapi/db/mgo"

	"github.com/globalsign/mgo/bson"
)

// Todo is todo
type Todo struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Topic  string        `json:"topic" bson:"topic"`
	Done   bool          `json:"done" bson:"done"`
	Status int           `json:"status" bson:"status"`
	Fake   bool          `json:"fake" bson:"fake"`
}

// CreateTodo is all todos
func CreateTodo(t *Todo) (*Todo, error) {
	var err error
	conn := mongo.MongoClient().Copy()
	defer conn.Close()

	err = conn.DB("document").C("todo").Insert(&t)
	return t, err
}

// UpdateTodo is all todos
func UpdateTodo(id bson.ObjectId, t *Todo) (*Todo, error) {
	var err error
	conn := mongo.MongoClient().Copy()
	defer conn.Close()

	err = conn.DB("document").C("todo").UpdateId(id, t)
	return t, err
}

// DeleteTodo is all todos
func DeleteTodo(id bson.ObjectId) error {
	var err error
	conn := mongo.MongoClient().Copy()
	defer conn.Close()

	err = conn.DB("document").C("todo").RemoveId(id)
	return err
}

// FindTodoByID is all todos
func FindTodoByID(id bson.ObjectId) (Todo, error) {

	var (
		todo Todo
		err  error
	)

	conn := mongo.MongoClient().Copy()
	defer conn.Close()

	err = conn.DB("document").C("todo").FindId(id).One(&todo)
	return todo, err
}

// FindAllTodos is all todos
func FindAllTodos(page int, limit int) ([]Todo, error) {

	var (
		todos []Todo
		err   error
	)

	skips := limit * (page - 1)

	conn := mongo.MongoClient().Copy()
	defer conn.Close()

	err = conn.DB("document").C("todo").Find(bson.M{}).Limit(limit).Skip(skips).Sort("-_id").All(&todos)
	return todos, err
}
