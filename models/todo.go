package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseName   = "test"
	collectionName = "test_result_2017"
)

type (

	//TodoStore is store JSON data
	TodoStore interface {
		GetTodo(id string) (Todo, error)
		GetAllTodo(page int64, limit int64) ([]*Todo, error)
		CreateTodo(*Todo) (*Todo, error)
		UpdateTodo(id string, todo *Todo) (*Todo, error)
		DeleteTodo(id string) (delete int64, err error)
	}

	//TodoModel is mongo db
	TodoModel struct {
		// Client
		//client *mongo.Client
		// Collections.
		//collection *mongo.Collection
		database *mongo.Database
	}

	// Todo is todo
	Todo struct {
		//ID primitive.ObjectID `json:"id" bson:"_id"`
		// ObjectId() or objectid.ObjectID is deprecated--use primitive instead
		ID primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		//ID interface{} `json:"id" bson:"_id"`
		//ID     primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		Topic     string    `json:"topic" bson:"topic"`
		Done      bool      `json:"done" bson:"done"`
		Status    int       `json:"status" bson:"status"`
		Fake      bool      `json:"fake" bson:"fake"`
		Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	}

	// TodoX is topic test
	TodoX struct {
		//ID    *primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		ID    primitive.ObjectID `json:"id" bson:"_id, omitempty"` // omitempty to protect against zeroed _id insertion
		Topic string             `json:"topic" bson:"topic"`
	}
)

// NewTodoModel is active new
/* func NewTodoModel(client *mongo.Client) *TodoModel {
	return &TodoModel{
		client: client,
	}
} */
func NewTodoModel(database *mongo.Database) *TodoModel {
	return &TodoModel{
		database: database,
	}
}

// CreateTodo is all todos
func (m *TodoModel) CreateTodo(todo *Todo) (*Todo, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("local time => ", time.Now().Local(), ", utc => ", time.Now().UTC())

	//time.Local = time.UTC

	todo.ID = primitive.NewObjectID()
	todo.Timestamp = time.Now().Local()
	todo.Done = false

	//res, err := m.client.Database("document").Collection("todo").InsertOne(ctx, &todo)
	res, err := m.database.Collection("todo").InsertOne(ctx, &todo)

	if err != nil {
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().Local().Format("2006-01-02T15:04:05Z"), err)
	} else {
		//fmt.Println("Inserted a Todo: ", res.InsertedID, res.InsertedID.(primitive.ObjectID).Hex())
		fmt.Println("Inserted a Todo: ", res.InsertedID.(primitive.ObjectID).Hex())
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().Local().Format("2006-01-02T15:04:05Z"), "test err")
	}

	return todo, err
}

// UpdateTodo is update todo by id
func (m *TodoModel) UpdateTodo(id string, todo *Todo) (*Todo, error) {
	var err error

	fmt.Println("update id:", id)

	// create ObjectID from string
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// set filters and updates
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": todo}

	fmt.Println("new update todo:", todo)

	//_, err = m.client.Database("document").Collection("todo").UpdateOne(ctx, filter, update)
	_, err = m.database.Collection("todo").UpdateOne(ctx, filter, update)
	//fmt.Println("Updated a Todo: ", res)

	return todo, err
}

// DeleteTodo is delete todo by id
func (m *TodoModel) DeleteTodo(id string) (int64, error) {
	var err error

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// delete document
	//res, err := m.client.Database("document").Collection("todo").DeleteOne(ctx, bson.M{"_id": objectID})
	res, err := m.database.Collection("todo").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}

	fmt.Printf("deleted count: %d\n", res.DeletedCount)
	// => deleted count: 1

	return res.DeletedCount, err
}

// GetTodo is get todo by id
func (m *TodoModel) GetTodo(id string) (Todo, error) {

	fmt.Println("call todo model time => ", time.Now().Local())

	var (
		todo Todo
		err  error
	)

	//fmt.Println("\n todo id type:", reflect.TypeOf(id))
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//log.Fatal(err)
		return todo, err
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	filter := bson.M{"_id": objectID}

	//err = m.client.Database("document").Collection("todo").FindOne(context.TODO(), filter).Decode(&todo)
	err = m.database.Collection("todo").FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		//log.Fatal(err)
		//timeStamp, _ := time.Parse("2006-01-02 15:04:05", todo.Timestamp.String())

		return todo, err
	}

	//fmt.Println("Timestamp => ", todo.Timestamp.String())

	//timeStamp, _ := time.Parse("2006-01-02T15:04:05.000Z", todo.Timestamp.String())
	timeStamp, _ := time.Parse("2006-01-02 15:04:05.000 +0000 UTC", todo.Timestamp.String())
	todo.Timestamp = timeStamp.Local()

	//fmt.Println("time => ", timeStamp.Format("2006-01-02 15:04:05"))

	fmt.Println("return data todo model time => ", time.Now().Local())

	//fmt.Printf("Found a single document: %+v\n", todo)
	return todo, err
}

// GetAllTodo is get todo by id
func (m *TodoModel) GetAllTodo(page int64, limit int64) ([]*Todo, error) {

	var (
		todos []*Todo
		err   error
	)

	skips := limit * (page - 1)

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//fmt.Printf("page: %v, limit: %v", page, limit)

	filter := bson.M{}
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skips)
	findOptions.SetSort(bson.M{"_id": -1})

	//cur, err := m.client.Database("document").Collection("todo").Find(context.TODO(), filter, findOptions)
	cur, err := m.database.Collection("todo").Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Todo
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Found all document: %+v\n", todos)
	return todos, err
}
