package main

import (
	"fmt"

	//"net/http"
	//"golangapi/db/mongo"
	"golangapi/db/elastics"
	"golangapi/db/mgo"
	"golangapi/db/mongo"
	"golangapi/router"

	//"golangapi/logger"
	//"golangapi/middlewares"
	//"golangapi/routers"

	//"golangapi/routers"
	"golangapi/handler"
	"strings"

	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	gormdb "golangapi/db/gorm"
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

	/* e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		//Skipper:      DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	})) */

	//start gorm db connect mysql
	gormdb.ConnectMySQL()

	//Start MongoDB Connect
	//Hold Mongo lib, It slower than mgo lib client
	mongo.ConnectMongo()

	//Start Mgo Connect
	mgo.ConnectMgo()

	//Start Elastics Connect
	elastics.ConnectES()

	// Start Router
	//routers.InitRoute(e)

	//Init Route
	//handler.InitialRoute()

	//Init Route
	router.InitialRoute(e)

	//Init Logs
	handler.InitialLogs(e)

	// Start Logger
	//middlewares.InitLog()

	/* fix test log */
	/*
		config := logger.Configuration{
			EnableConsole:     true,
			ConsoleLevel:      logger.Debug,
			ConsoleJSONFormat: true,
			EnableFile:        true,
			FileLevel:         logger.Info,
			FileJSONFormat:    true,
			FileLocation:      "log.log",
		}
		err := logger.NewLogger(config, logger.InstanceZapLogger)
		if err != nil {
			log.Fatalf("Could not instantiate log %s", err.Error())
		}

		contextLogger := logger.WithFields(logger.Fields{"key1": "value1"})
		contextLogger.Debugf("Starting with zap")
		contextLogger.Infof("Zap is awesome")

		err = logger.NewLogger(config, logger.InstanceLogrusLogger)
		if err != nil {
			log.Fatalf("Could not instantiate log %s", err.Error())
		}
		contextLogger = logger.WithFields(logger.Fields{"key1": "value1"})
		contextLogger.Debugf("Starting with logrus")

		contextLogger.Infof("Logrus is awesome")
	*/
	/* end fix test log */

	port := fmt.Sprintf(":%v", viper.GetString("port"))
	e.Logger.Fatal(e.Start(port))
}
