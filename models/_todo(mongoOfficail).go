package models

import (
	"golangapi/db/mongo"
	//"golangapi/db/mgo"

	"fmt"
	"time"
	"context"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	//"go.mongodb.org/mongo-driver/mongo"
    //"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	
	/*
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	*/

	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/globalsign/mgo/bson"
)

// Todo is todo
type Todo struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	//ID    bson.ObjectId `json:"id" bson:"_id"`
	Topic string        `json:"topic" bson:"topic"`
	Done  bool          `json:"done" bson:"done"`
}

// CreateTodo is all todos
func CreateTodo(t *Todo) (*Todo, error) {
	var err error
	client := mongo.ClientManager()
	
	// create a new timeout context
	/*
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create a mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// disconnect from mongo
	defer client.Disconnect(ctx)

	// select collection from database
	col := client.Database("document").Collection("todo")

	//t.ID = primitive.NewObjectID()
	res, err := col.InsertOne(ctx, &t)
	if(err != nil){
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), err)
	}

	fmt.Println("Inserted a Logger: ", res.InsertedID)
	*/
	//fmt.Println("Inserted a Logger: ", res.InsertedID, res.InsertedID.(primitive.ObjectID).Hex())

	//client := mgo.MongoManager()
	//defer client.Close()

	//fmt.Println("Inserted a Todo, client : ", client)

	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//err = client.Database("document").Collection("todo").InsertOne(contact.TODO(), &t)
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//res, err := client.Database("document").Collection("todo").InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()

	//fmt.Println("Inserted a Todo, ctx : ", ctx)
	/*
	newTodo := bson.M{
		"topic": t.Topic,
		"done": t.Done,
	}
	*/
	//res, err := client.Database("document").Collection("todo").InsertOne(ctx, newTodo)
	
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.ID = primitive.NewObjectID()
	res, err := client.Database("document").Collection("todo").InsertOne(ctx, &t)
	

	
	//t.ID = bson.NewObjectId()
	//err = client.DB("document").C("todo").Insert(&t)

	if(err != nil){
		//fmt.Println("\n err Create Todo => ", err)
		//logger.SetOutput(&middlewares.Logger{Collection: "logger"}) //middleware log to mongodb
		
		log.Printf(`{"time": "%s", "message": "{%s}", "level": "error"`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), err)
	} else {
		//idx := res.InsertedID.(primitive.ObjectID)
		//newTodo := t
		//newTodo.ID = idx
		//t.ID =  idx
		
		fmt.Println("Inserted a Todo: ", res.InsertedID, res.InsertedID.(primitive.ObjectID).Hex())
	}

	return t, err
}

// UpdateTodo is all todos
func UpdateTodo(idx string, t *Todo) (*Todo, error) {
	var err error
	client := mongo.ClientManager()

	// create ObjectID from string
	id, err := primitive.ObjectIDFromHex(idx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// set filters and updates
	filter := bson.M{"_id": id}
	update := bson.M{"$set": t}

	fmt.Println("new update todo:", t)

	res, err := client.Database("document").Collection("todo").UpdateOne(ctx, filter, update)
	fmt.Println("Updated a Todo: ", res)
	return t, err
}

// DeleteTodo is all todos
func DeleteTodo(idx string) error {
	var err error
	client := mongo.ClientManager()
	
	// create ObjectID from string
	id, err := primitive.ObjectIDFromHex(idx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// delete document
	res, err := client.Database("document").Collection("todo").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted count: %d\n", res.DeletedCount)
	// => deleted count: 1

	return err
}

// FindTodoByID is all todos
func FindTodoByID(idx string) (Todo, error) {

	var (
		todo Todo
		err  error
	)

	// create ObjectID from string
	id, err := primitive.ObjectIDFromHex(idx)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// filter posts tagged as golang
	//filter := bson.M{"tags": bson.M{"$elemMatch": bson.M{"$eq": "golang"}}}
	filter := bson.M{"_id": id}

	client := mongo.ClientManager()
	if err := client.Database("document").Collection("todo").FindOne(ctx, filter).Decode(&todo); err != nil {
		log.Fatal(err)
	}

	return todo, err
}

// FindAllTodos is all todos
func FindAllTodos() ([]Todo, error) {

	var (
		todos []Todo
		todo Todo
		err   error
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}

	client := mongo.ClientManager()

	// find all documents
	cursor, err := client.Database("document").Collection("todo").Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("cursor: %+v\n", cursor)

	// iterate through all documents
	for cursor.Next(ctx) {
		// decode the document
		//var t Todo
		if err := cursor.Decode(&todo); err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
		fmt.Printf("todo: %+v\n", todo)
	}

	// check if the cursor encountered any errors while iterating 
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return todos, err
}