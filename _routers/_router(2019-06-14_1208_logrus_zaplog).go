package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//"net/http/httptrace"
)

// TimeEncoder return time encode
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02T15:04:05Z07:00"))
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z"))
}

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
				logger.Printf(`{"time": "%s", "prefix": "Logger", "message": "{}", "level": "info","data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
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

	// Log as JSON instead of the default ASCII formatter.
	/*logrus.SetFormatter(&logrus.JSONFormatter{
		// FieldMap: logrus.FieldMap{
		// 	logrus.FieldKeyTime:  "@timestamp",
		// 	logrus.FieldKeyLevel: "@level",
		// 	logrus.FieldKeyMsg: "@message",
		// 	logrus.FieldKeyFunc:  "@caller",
		// },
	})*/

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)
	//logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	//ljack := &middlewares.LoggerLumberjack()
	//mWriter := io.MultiWriter(os.Stderr, &middlewares.Logrus{Collection: "logger"})
	//log.SetOutput(mWriter)
	//logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	mWriter := io.MultiWriter(os.Stderr, &middlewares.Logrus{Collection: "logger"})
	logrus.SetOutput(mWriter)

	// Only log the warning severity or above.
	//logrus.SetLevel(logrus.WarnLevel)
	logrus.SetLevel(logrus.InfoLevel)

	// Caller func report
	logrus.SetReportCaller(true)

	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Printf("==> %s\n", resBody)
	// }))

	// w := zapcore.AddSync(&lumberjack.Logger{
	//     Filename:   "foo.log",
	//     MaxSize:    500, // megabytes
	//     MaxBackups: 3,
	//     MaxAge:     28, // days
	// })

	//hook := zapcore.AddSync(&middlewares.Logrus{Collection: "logger"})

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:   "time",
		LevelKey:  "level",
		NameKey:   "logger",
		CallerKey: "caller",
		//MessageKey:    "msg",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		//EncodeTime:    zapcore.ISO8601TimeEncoder,
		//EncodeTime:     zapcore.TimeEncoder(zapcore.PrimitiveArrayEncoder.AppendString(time.Time.Format("2006-01-02 15:04:05.000"))),
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		//zapcore.NewConsoleEncoder(NewEncoderConfig()),
		zapcore.NewJSONEncoder(encoderConfig),
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), hook),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&middlewares.Logrus{Collection: "logger"})),
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(&middlewares.Logrus{Collection: "logger"})),
		//zap.DebugLevel,
		zap.InfoLevel,
	)

	zaplogger := zap.New(core, zap.AddCaller())
	//zaplogger.Info("info test xxx")
	zaplogger.Info("info test xxx", zap.String("message", "{}"))

	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {

			// req, _ := http.NewRequest("GET", "http://sanook.com", nil)
			// //req := c.Request()

			// fmt.Println("req => ", req)

			// trace := &httptrace.ClientTrace{
			// 	GotConn: func(connInfo httptrace.GotConnInfo) {
			// 		fmt.Printf("Got Conn: %+v\n", connInfo)
			// 	},
			// 	DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			// 		fmt.Printf("DNS Info: %+v\n", dnsInfo)
			// 	},
			// }
			// req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
			// _, err := http.DefaultTransport.RoundTrip(req)

			// //req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
			// //_, err := http.DefaultTransport.RoundTrip(req)
			// if err != nil {
			// 	//fmt.Println("err => ", err)
			// 	log.Fatal(err)
			// }

			reqB := `""`
			if len(reqBody) > 0 {
				reqB = string(reqBody)
			}

			jsonStr := fmt.Sprintf(`{"time": "%s", "prefix": "API", "message": "{}", "data": {"id":"%s","req":%s,"res":%s}}`, time.Now().UTC().Format("2006-01-02T15:04:05Z"), c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			//fmt.Println("jsonStr => ", jsonStr)

			var jsonData map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
				panic(err)
			}

			logrus.WithFields(jsonData).Info("tracked logs request/response from logrus")

			//zaplogger.Info("info test xxx", zap.String("message", "{}"))
			jsonStr2 := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
			var jsonData2 map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr2), &jsonData2); err != nil {
				panic(err)
			}

			//zaplogger.Info("Test Zap Logger")
			zaplogger.With(
				//zap.Namespace("data"),
				//zap.String("id", c.Response().Header().Get(echo.HeaderXRequestID)),
				//zap.String("req", reqB),
				//zap.String("res", string(resBody)),
				zap.String("message", "{}"),
				zap.Any("data", jsonData2),
			).Info("tracked logs request/response from zap")
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
