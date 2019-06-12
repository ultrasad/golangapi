package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golangapi/controllers"
	"golangapi/middlewares"

	//"strconv"
	//"github.com/google/logger"

	//"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

	/* Uberzap OK, working for logfile and stdout
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

	// Log as JSON instead of the default ASCII formatter.
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

	// Caller func report
	logrus.SetReportCaller(true)

	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Printf("==> %s\n", resBody)
	// }))

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {

			reqB := `""`
			if len(reqBody) > 0 {
				reqB = string(reqBody)
			}

			jsonStr := fmt.Sprintf(`{"time": "%s", "prefix": "API", message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)

			fmt.Println("jsonStr => ", jsonStr)

			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
				panic(err)
			}

			logrus.WithFields(jsonData).Info("API Request")
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
