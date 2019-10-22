package main

import (
	"fmt"
	//"net/http"
	//"golangapi/db/mongo"
	"golangapi/db/elastics"
	"golangapi/db/mgo"
	"golangapi/middlewares"
	"golangapi/routers"

	//"golangapi/routers"
	"golangapi/handler"
	"strings"

	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
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
	//Hold Mongo lib, It slower than mgo lib client
	//mongo.ConnectMongo()

	//Start Mgo Connect
	mgo.ConnectMgo()

	//Start Elastics Connect
	elastics.ConnectES()

	// Start Router
	routers.InitRoute(e)

	//Init Route
	//handler.InitialRoute()

	//Init Route
	handler.InitialRoute(e)

	//Init Logs
	handler.InitialLogs(e)

	// Start Logger
	middlewares.InitLog()

	port := fmt.Sprintf(":%v", viper.GetString("port"))
	e.Logger.Fatal(e.Start(port))
}
