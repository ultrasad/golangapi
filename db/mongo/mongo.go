package mongo

import (
	"fmt"
	"log"
	//"time"
	"context"
	//"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/spf13/viper"
)

var (
	mongoDBClient *mongo.Client
	err   error
)

// ConnectMongo return Mongo Connection
func ConnectMongo() *mongo.Client {
	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	ctx := context.Background()
	connString := fmt.Sprintf("%v:%v@%v", mongoUser, mongoPass, mongoHost)
	clientOptions := options.Client().ApplyURI("mongodb://" + connString)
	mongoDBClient, err = mongo.Connect(ctx, clientOptions)

	/*
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//db := client.Database(mongoDB)
	*/

	if err != nil {
        log.Fatal(err)
    }

    err = mongoDBClient.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

	return mongoDBClient

	//connString := fmt.Sprintf("%v:%v@%v", mongoUser, mongoPass, mongoHost)

	// create a new context
	/*
	ctx := context.Background()

	// create a mongo client
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://localhost:6548/"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// connect to mongo
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// disconnects from mongo
	defer client.Disconnect(ctx)

	fmt.Println("Connected to MongoDB!")

	return client
	*/
}

// ClientManager return MongoDB Session
func ClientManager() *mongo.Client {
	fmt.Println("Call Client Manager... ")
	return mongoDBClient
}

/*
var (
	mgoDB *mgo.Session
	err   error
)

// ConnectMgo MongoDB Connect
func ConnectMgo() *mgo.Session {
	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	connString := fmt.Sprintf("%v:%v@%v", mongoUser, mongoPass, mongoHost)
	mgoDB, err = mgo.Dial(connString)
	if err != nil {
		log.Printf("dial mongodb server with connection string %q: %v", connString, err)
		//return
	}

	return mgoDB
}

// MgoManager return MongoDB Session
func MgoManager() *mgo.Session {
	return mgoDB
}
*/