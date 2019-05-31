package routers

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golangapi/controllers"
	"golangapi/middlewares"

	//"strconv"

	//"github.com/google/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/spf13/viper"
	//li "workshop01/middlewares/logwrapper"
	//zaplogger "workshop01/middlewares/uberzap"
)

//Init func
func Init(e *echo.Echo) {

	//OK
	//middlewares.Init()
	//log.SetPrefix("prefix")
	//2019-05-14, change log to logrus
	//2019-05-17, move to Init
	
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
			/*
			if(resBody != nil){
				logger.Printf(`{"time": "%s", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			} else {
				fmt.Println("response nil")
			}*/
		},
	}))

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "hanajung" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
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
