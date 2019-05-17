package mgo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

var (
	mongoDB *mgo.Session
	err   error
)

// ConnectMgo MongoDB Connect
func ConnectMgo() *mgo.Session {
	mongoHost := viper.GetString("mongo.host")
	mongoUser := viper.GetString("mongo.user")
	mongoPass := viper.GetString("mongo.pass")

	connString := fmt.Sprintf("%v:%v@%v", mongoUser, mongoPass, mongoHost)
	mongoDB, err = mgo.Dial(connString)
	if err != nil {
		log.Printf("dial mongodb server with connection string %q: %v", connString, err)
		//return
	}

	return mongoDB
}

// MongoManager return MongoDB Session
func MongoManager() *mgo.Session {
	return mongoDB
}