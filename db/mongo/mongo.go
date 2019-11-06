package mongo

import (
	"context"
	"fmt"
	"log"
	"sync"
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

	//DBNewDataStoreClient test new client
	DBNewDataStoreClient *mongo.Client

	err error
)

//CONNECTED ...
const CONNECTED = "Successfully connected to database: %v"

type (
	//DBNewDataStore ...
	DBNewDataStore struct {
		db      *mongo.Database
		Session *mongo.Client
		//logger  *logrus.Logger
	}
)

//connectSingleton is check connect
func connectSingleton() *mongo.Client {
	if mongoDBClient == nil {
		fmt.Println("mongo db lose connect...")
		mongoDBClient = ConnectMongo()
	}
	return mongoDBClient
}

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

	//var connectOnce sync.Once

	//client, err = mongo.Connect(ctx, options.Client().ApplyURI(connString))
	//connectOnce.Do(func() {

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

	//})

	//new mongodb connect solution
	//DataStoreNew()

	return mongoDBClient
}

// ClientManager return MongoDB Session
/* func ClientManager() *mongo.Client {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
} */
/* func ClientManager() *mongo.Client {
	//fmt.Println("Call Mongo Client Manager.")
	return mongoDBClient
} */
func ClientManager() *mongo.Client {
	if mongoDBClient == nil {
		fmt.Println("mongo db lose connect...")
		mongoDBClient = ConnectMongo()
	}

	return mongoDBClient
}

//DataStoreNew is new mongodb data store
func DataStoreNew() *DBNewDataStore {

	var mongoDataStore *DBNewDataStore
	//mongoDataStore := new(DBNewDataStore)
	db, session := connect()
	if db != nil && session != nil {

		// log statements here as well
		fmt.Println("mongo db and session.")

		mongoDataStore = new(DBNewDataStore)
		mongoDataStore.db = db
		//mongoDataStore.logger = logger
		mongoDataStore.Session = session

		//new client
		//DBNewDataStoreClient = session
		return mongoDataStore
	}

	//logger.Fatalf("Failed to connect to database: %v", config.DatabaseName)
	fmt.Println("Failed to connect to database.")

	return nil
}

func connect() (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		fmt.Println("Mongo connectOnce.")
		db, session = connectToMongo()
	})

	return db, session
}

//Dconnect ...
/* func Dconnect() (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		fmt.Println("Mongo connectOnce.")
		db, session = connectToMongo()

		//fix test
		DBNewDataStoreClient = session
	})

	return db, session
} */

func connectToMongo() (a *mongo.Database, b *mongo.Client) {

	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	connString := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPass, mongoHost)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	session, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		//logger.Fatal(err)
		fmt.Println("Mongo clinet err: ", err)
	}
	session.Connect(context.TODO())
	if err != nil {
		//logger.Fatal(err)
		fmt.Println("Mongo connect err: ", err)
	}

	err = session.Ping(ctx, nil)

	if err != nil {
		//fmt.Println("ping mongodb err,", err)
		log.Fatal(err)
	}

	var DB = session.Database("document")
	//logger.Info(CONNECTED, generalConfig.DatabaseName)
	fmt.Println("Mongo NewDatastore CONNECTED.")

	return DB, session
}

//ClientNewManager return new db
func ClientNewManager() *mongo.Client {
	return DBNewDataStoreClient
}
