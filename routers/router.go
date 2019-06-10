package routers

import (
	//"fmt"

	"net/http"

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

	logrus "github.com/sirupsen/logrus"
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
	//jsonStr := `{"id": "", "req": "Test Req", "res": "Test Res"}`
	logrus.SetFormatter(&logrus.JSONFormatter{
		/*FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name", // non-ECS
		},*/
		FieldMap: logrus.FieldMap{
			//logrus.FieldKeyMsg: "message",
		},
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	//logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	//ljack := &middlewares.LoggerLumberjack()
	//mWriter := io.MultiWriter(os.Stderr, &middlewares.Logrus{Collection: "logger"})
	//log.SetOutput(mWriter)
	logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {

			reqB := "\"\""
			if len(reqBody) > 0 {
				reqB = string(reqBody)
			}

			//logger.Printf(`{"time": "%s", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)

			//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//var val []byte = []byte(fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody))
			//s, _ := strconv.Unquote(string(val))
			//jsonStr := &middlewares.Logrus{Data: s}

			//jsonStr := []string{"1", "2", "3"}
			logrus.WithFields(logrus.Fields{
				//"type": "Animal",
				//"name": "Chuche",
				//"data":   {"id": "", "req": reqB, "res": reqBody},
				//"data": fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody),
				//"data":    &middlewares.Logrus{Data: {}},
				"message": "{}",
				//"id":      c.Response().Header().Get(echo.HeaderXRequestID),
				//"req":     reqB,
				//"res":     resBody,
				"data": &middlewares.CtxLogger{ID: c.Response().Header().Get(echo.HeaderXRequestID), Req: reqB, Res: resBody},
			}).Warning("A walrus appears")

			/*
				logrus.WithFields(logrus.Fields{
					"animal": "walrus",
					"size":   10,
				}).Warning("A group of walrus emerges from the ocean")
			*/

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
