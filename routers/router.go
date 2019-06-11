package routers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"

	"golangapi/controllers"
	"golangapi/middlewares"

	//"strconv"
	//"github.com/google/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//logrus "github.com/sirupsen/logrus"
	//logrus "github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
	//li "workshop01/middlewares/logwrapper"
	//zaplogger "workshop01/middlewares/uberzap"
	//"github.com/labstack/gommon/log"
	//logrus "github.com/sirupsen/logrus"
	logrus "github.com/sirupsen/logrus"
	//zaplogger "golangapi/middlewares/uberzap"
	//uberzap "go.uber.org/zap"
	//uberzap "go.uber.org/zap"
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

				reqB := `""`
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
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	//ljack := &middlewares.LoggerLumberjack()
	//mWriter := io.MultiWriter(os.Stderr, &middlewares.Logrus{Collection: "logger"})
	//log.SetOutput(mWriter)
	//logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)

	logrus.SetReportCaller(true)

	/*
		config := uberzap.NewProductionConfig()

		config.OutputPaths = []string{
			//callback(),
			"stdout",
		}

		config.EncoderConfig.LevelKey = "level"
		//config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.CallerKey = "caller"
		config.EncoderConfig.MessageKey = "message"

		logger, err := config.Build()
		if err != nil {
			panic(err)
		}
		defer logger.Sync()
	*/
	//logger.Info("logger construction succeeded")

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {

			//reqB := "\"\""
			// reqB := `""`
			//reqB := []byte(``)
			/*if len(reqBody) < 1 {
				//reqBody = []byte(`""`)
				//reqB = reqBody
				reqBody = []byte(`{}`)
				//fmt.Println("req empty...")
				//reqB = []byte(reqBody)
			}*/

			reqB := `""`
			if len(reqBody) > 0 {
				//reqBody = []byte(`""`)
				reqB = string(reqBody)
				//reqBody = []byte(`{}`)
				//fmt.Println("req empty...")
				//reqB = []byte(reqBody)
			}

			//fmt.Println("reqBody => ", reqBody)
			//logger.Info("logger construction succeeded")

			//fmt.Println("req => ", reqBody)
			//fmt.Println("res => ", resBody)

			//logger.Printf(`{"time": "%s", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)

			//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//var val []byte = []byte(fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody))
			//s, _ := strconv.Unquote(string(val))

			//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, "1", "{}", "{}")
			//jsonStr := fmt.Sprintf(`{"time": "%s", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			// jsonStr := fmt.Sprintf(`{data": {"id":"%s","request":%s,"response":%s}}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//byt := []byte(jsonStr)
			// //fmt.Println(jsonStr)

			jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			var newData map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr), &newData); err != nil {
				panic(err)
			}
			
			//
			fmt.Println("newData", newData)

			jsonData := map[string]interface{}{
				"message": "{}",
				"time":    time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				"data": newData,
			}

			logrus.WithFields(jsonData).Warning("Test Message by Hanajung !!!")

			/*
			var reqB map[string]interface{}
			//var reqB []interface{}
			if err := json.Unmarshal(reqBody, &reqB); err != nil {
				panic(err)
			}
			//fmt.Println(reqB)

			var resB map[string]interface{} //object
			//var resB []interface{} //array
			if err := json.Unmarshal(resBody, &resB); err != nil {
				panic(err)
			}
			//fmt.Println(resB)

			jsonData := map[string]interface{}{
				"message": "{}",
				"time":    time.Now().UTC().Format("2006-01-02T15:04:05Z"),
				"data": map[string]interface{}{
					"id":  c.Response().Header().Get(echo.HeaderXRequestID),
					"req": reqB,
					"res": resB,
				},
			}
			*/

			//jsonData["data"] = jsonStr

			//jsonStr := []byte(`{"id": "1"}`)

			//jsonStr := []string{"1", "2", "3"}
			/*jsonData := map[string]interface{}{
				"message": "{}",
				"data": map[string]interface{}{
					"id": c.Response().Header().Get(echo.HeaderXRequestID),
					"req": map[string]interface{}{
						"test": "xxx",
					},
					//"req": reqB,
					//"req": map[string]interface{}{
					//	"test": "xxx",
					//},
				},
				//"res": resBody,
			}*/

			/*jsonData := &middlewares.CtxLogger{
				ID:  c.Response().Header().Get(echo.HeaderXRequestID),
				Req: "{test:xxx}",
				//"res": resBody,
			}*/

			//Simple Employee JSON object which we will parse
			/*empJSON := `{
				"id": 11,
				"message": "Hanajung",
				"data": {
					"id": "Mumbai",
					"req": "{\"test\":\"xxx\"}",
					"res": "{}"
				}
			}`

			// Declared an empty interface
			var result map[string]interface{}

			// Unmarshal or Decode the JSON to the interface.
			json.Unmarshal([]byte(empJSON), &result)*/

			//myData := make(map[string]interface{})
			//myData["req"] = reqB
			//myData["res"] = resBody
			/*
				jsonbody, err := json.Marshal(myData)
				if err != nil {
					// do error check
					fmt.Println(err)
					return
				}
			*/

			/*
				logrus.WithFields(logrus.Fields{
					//"type": "Animal",
					//"name": "Chuche",
					//"data":   {"id": "", "req": reqB, "res": reqBody},
					//"data": fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody),
					//"data":    &middlewares.Logrus{Data: {}},
					//"message": "{}",
					//"id":      c.Response().Header().Get(echo.HeaderXRequestID),
					//"req":     reqB,
					//"res":     resBody,
					//"data": &middlewares.CtxLogger{ID: c.Response().Header().Get(echo.HeaderXRequestID), Req: reqB, Res: resBody},
					//"data": {"id": "1"},
					//"prefix": "logs prefix",
					"data": jsonData,
				}).Warning("A walrus appears")
			*/

			//logrus.WithFields(jsonData).Warning("Test Message by Hanajung !!!")

			/*log.SetPrefix("api")
			logger := log.New(os.Stderr, "", 0)
			logger.SetOutput(&middlewares.Logrus{Collection: "logger"}) //middleware log to mongodb
			logger.Printf(jsonStr)*/
		},
	}))

	//Uber zap logger, 2019-06-11, It work
	/*
		zaplog, loggerError := zaplogger.NewLogger("zaplogs", func() string {
			year, month, day := time.Now().Date()
			//return "logs/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + ".json"
			return "zaplogs/" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + ".log"
		})

		if loggerError != nil {
			e.Logger.Fatal(loggerError)
		}

		defer zaplog.Sync()
		zaplog.Info("First Load")
	*/

	/*
		config := uberzap.NewProductionConfig()

		config.OutputPaths = []string{
			//callback(),
			"stdout",
		}

		config.EncoderConfig.LevelKey = "level"
		//config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.TimeKey = ""
		config.EncoderConfig.CallerKey = "caller"
		config.EncoderConfig.MessageKey = "message"

		logger, err := config.Build()
		if err != nil {
			panic(err)
		}
		defer logger.Sync()

		logger.Info("logger construction succeeded")
	*/

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
