package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//MongoDBClient ...
	//mongoDBClient *mongo.Client

	//mongoDBCollection ...
	mongoDBClient *mongo.Database

	err error
)

// ConnectMongo return Mongo Connection
//func ConnectMongo() *mongo.Client {
func ConnectMongo() *mongo.Database {
	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	ctx := context.TODO()
	connString := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPass, mongoHost)

	clientOptions := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	mongoDBClient = client.Database("document")

	fmt.Println("Connected to MongoDB!")

	return mongoDBClient
}

// ClientManager return MongoDB Session
/* func ClientManager() *mongo.Client {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
} */
func ClientManager() *mongo.Database {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
}
