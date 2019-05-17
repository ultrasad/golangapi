package main

import (
	"fmt"
	//"net/http"
	"golangapi/db/mongo"
	"golangapi/db/mgo"
	"golangapi/routers"
	"strings"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("port", "8083")

	e := echo.New()

	//Start MongoDB Connect
	mongo.ConnectMongo()

	//Start Mgo Connect
	mgo.ConnectMgo()

	// Start Router
	routers.Init(e)

	port := fmt.Sprintf(":%v", viper.GetString("port"))
	e.Logger.Fatal(e.Start(port))
}
