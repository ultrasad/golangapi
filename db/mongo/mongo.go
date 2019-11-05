package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//MongoDBClient ...
	mongoDBClient *mongo.Client

	//mongoDBCollection ...
	//mongoDBClient *mongo.Database

	err error
)

// ConnectMongo return Mongo Connection
func ConnectMongo() *mongo.Client {
	//func ConnectMongo() *mongo.Database {
	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//mongoDBClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	//ctx := context.TODO()
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connString := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPass, mongoHost)

	//clientOptions := options.Client().ApplyURI(connString)
	//mongoDBClient, err := mongo.Connect(ctx, clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//client, err = mongo.Connect(ctx, options.Client().ApplyURI(connString))
	mongoDBClient, err = mongo.NewClient(options.Client().ApplyURI(connString))
	err = mongoDBClient.Connect(ctx)

	if err != nil {
		//fmt.Println("connect mongodb err,", err)
		log.Fatal(err)
	}

	err = mongoDBClient.Ping(ctx, nil)

	if err != nil {
		//fmt.Println("ping mongodb err,", err)
		log.Fatal(err)
	}

	//mongoDBClient = client.Database("document")

	fmt.Println("Connected to MongoDB!")

	return mongoDBClient
}

// ClientManager return MongoDB Session
/* func ClientManager() *mongo.Client {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
} */
func ClientManager() *mongo.Client {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
}
