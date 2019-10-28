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

type (

	//TodoStore is store JSON data
	TodoStore interface {
		GetTodo(id string) (Todo, error)
		GetAllTodo(page int64, limit int64) ([]*Todo, error)
		CreateTodo(*Todo) (*Todo, error)
		UpdateTodo(id string, todo *Todo) (*Todo, error)
	}

	//TodoModel is mongo db
	TodoModel struct {
		client *mongo.Client
	}

	// Todo is todo
	Todo struct {
		//ID primitive.ObjectID `json:"id" bson:"_id"`
		// ObjectId() or objectid.ObjectID is deprecated--use primitive instead
		ID primitive.ObjectID `bson:"_id, omitempty"`
		//ID interface{} `json:"id" bson:"_id"`
		//ID     primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		Topic  string `json:"topic" bson:"topic"`
		Done   bool   `json:"done" bson:"done"`
		Status int    `json:"status" bson:"status"`
		Fake   bool   `json:"fake" bson:"fake"`
	}

	// TodoX is topic test
	/* TodoX struct {
		//ID    *primitive.ObjectID `json:"id" bson:"_id, omitempty"`
		IDX   interface{} `json:"idx" bson:"_id, omitempty"` // omitempty to protect against zeroed _id insertion
		Topic string      `json:"topic" bson:"topic"`
	} */
)

// NewTodoModel is active new
func NewTodoModel(client *mongo.Client) *TodoModel {
	return &TodoModel{
		client: client,
	}
}

// CreateTodo is all todos
func (m *TodoModel) CreateTodo(todo *Todo) (*Todo, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	todo.ID = primitive.NewObjectID()
	todo.Done = false

	res, err := m.client.Database("document").Collection("todo").InsertOne(ctx, &todo)

	if err != nil {
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), err)
	} else {
		//fmt.Println("Inserted a Todo: ", res.InsertedID, res.InsertedID.(primitive.ObjectID).Hex())
		fmt.Println("Inserted a Todo: ", res.InsertedID.(primitive.ObjectID).Hex())
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), "test err")
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

	res, err := m.client.Database("document").Collection("todo").UpdateOne(ctx, filter, update)
	fmt.Println("Updated a Todo: ", res)

	return todo, err
}

// GetTodo is get todo by id
func (m *TodoModel) GetTodo(id string) (Todo, error) {

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

	err = m.client.Database("document").Collection("todo").FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		//log.Fatal(err)
		return todo, err
	}

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

	cur, err := m.client.Database("document").Collection("todo").Find(context.TODO(), filter, findOptions)
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
