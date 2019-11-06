package _mongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CONNECTED ...
const CONNECTED = "Successfully connected to database: %v"

type (
	//DBNewDatastore ...
	DBNewDatastore struct {
		db      *mongo.Database
		Session *mongo.Client
		//logger  *logrus.Logger
	}
)

//DataStoreNew is new mongodb data store
func DataStoreNew() *DBNewDatastore {

	var mongoDataStore *DBNewDatastore
	db, session := connect()
	if db != nil && session != nil {

		// log statements here as well

		mongoDataStore = new(DBNewDatastore)
		mongoDataStore.db = db
		//mongoDataStore.logger = logger
		mongoDataStore.Session = session
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
		fmt.Println("Mongo connectOnce...")
		db, session = connectToMongo()
	})

	return db, session
}

func connectToMongo() (a *mongo.Database, b *mongo.Client) {

	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	connString := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPass, mongoHost)

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

	var DB = session.Database("document")
	//logger.Info(CONNECTED, generalConfig.DatabaseName)
	fmt.Println("Mongo NewDatastore CONNECTED.")

	return DB, session
}
