package routers

import (
	//"fmt"

	"net/http"
	"os"

	"golangapi/controllers"
	"golangapi/middlewares"

	//"strconv"
	//"github.com/google/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//"github.com/spf13/viper"
	//li "workshop01/middlewares/logwrapper"
	//zaplogger "workshop01/middlewares/uberzap"
	//"github.com/labstack/gommon/log"

	log "github.com/sirupsen/logrus"
)

//Init func
func Init(e *echo.Echo) {

	//OK
	//middlewares.Init()
	//log.SetPrefix("prefix")
	//2019-05-14, change log to logrus
	//2019-05-17, move to Init

	//Old ok, master 2019-05-31
	/*
		log.SetPrefix("api")
		logger := log.New(os.Stderr, "", 0)
		logger.SetOutput(&middlewares.Logger{Collection: "logger"}) //middleware log to mongodb

		// Logger request, response
		//logs fmt
		e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
			Handler: func(c echo.Context, reqBody, resBody []byte) {

				//fmt.Println("req => ", reqBody)
				//fmt.Println("res => ", resBody)

				reqB := "\"\""
				if len(reqBody) > 0 {
					reqB = string(reqBody)
				}

				//fmt.Println(reqB)
				logger.Printf(`{"time": "%s", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			},
		}))
	*/

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {

			log.WithFields(log.Fields{
				"animal": "walrus",
			}).Warning("A walrus appears")

			log.WithFields(log.Fields{
				"animal": "walrus",
				"size":   10,
			}).Warning("A group of walrus emerges from the ocean")

		},
	}))

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "hanajung" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"level":"${level}", "time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\r\n",
		//CustomTimeFormat: "2006-01-02T15:04:05Z",
		Output: &middlewares.Logs{Collection: "logs"},
		//Output: os.Stdout,
		//Output: echoLog,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/todos", controllers.List)
	e.POST("/todos", controllers.Create)
	e.GET("/todos/:id", controllers.View)
	e.PUT("/todos/:id", controllers.Done)
	e.DELETE("/todos/:id", controllers.Delete)

	e.GET("/allusers", controllers.GetAllUser)

	//GoRoutine
	e.GET("/hello", controllers.CallHelloRoutine)

	//Elastics Route
	//e.GET("/esversion", controllers.ESVersion)

	//Elastics Search
	e.GET("/essearch", controllers.Search)

	//e.Logger.Fatal(e.Start(port))
}
