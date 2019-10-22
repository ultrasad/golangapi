package handler

import (
	"encoding/json"
	"fmt"
	"golangapi/db/mgo"
	"os"
	"path/filepath"

	//"golangapi/logger"
	//"golangapi/middlewares"

	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"gopkg.in/natefinch/lumberjack.v2"
)

type (
	// Logs struct log from echo
	Logs struct {
		//ID           string    `json:"id" bson:"id"`
		//ID           primitive.ObjectID    `json:"id" bson:"_id"`
		Time         time.Time `bson:"time" json:"time"`
		RemoteIP     string    `bson:"remote_ip" json:"remote_ip"`
		Host         string    `bson:"host" json:"host"`
		Method       string    `bson:"method" json:"method"`
		URI          string    `bson:"uri" json:"uri"`
		Status       int       `bson:"status" json:"status"`
		Latency      int       `bson:"latency" json:"latency"`
		LatencyHuman string    `bson:"latency_human" json:"latency_human"`
		BytesIn      int       `bson:"bytes_in" json:"bytes_in"`
		BytesOut     int       `bson:"bytes_out" json:"bytes_out"`
		Collection   string    `bson:"-"`
	}

	//CtxLogger struct logger req,res
	CtxLogger struct {
		ID  string      `bson:"id" json:"id"`
		Req interface{} `bson:"request" json:"req"`
		Res interface{} `bson:"response" json:"res"`
	}

	//Zaplog struct log from Zap logger
	Zaplog struct {
		ID       string `json:"id" bson:"id"`
		RemoteIP string `bson:"remote_ip" json:"remote_ip"`
		Host     string `bson:"host" json:"host"`
		Method   string `bson:"method" json:"method"`
		URI      string `bson:"uri" json:"uri"`
		Status   int    `bson:"status" json:"status"`
		Latency  string `bson:"latency" json:"latency"`

		Time       time.Time `bson:"time" json:"time"`
		Lv         string    `bson:"level" json:"level"`
		Prefix     string    `bson:"prefix" json:"prefix"`
		Message    string    `bson:"-" json:"message"`
		Msg        string    `bson:"msg" json:"msg"`
		Data       CtxLogger `bson:"data" json:"data"`
		Collection string    `bson:"-"`
	}
)

//echo Logs
// 2019-05-14, comment fix test
func (lg *Logs) Write(logEcho []byte) (n int, err error) {

	err = json.Unmarshal(logEcho, &lg)
	if err != nil {
		fmt.Println("\n err Logs, json Unmarshal >>>", err)
		return
	}

	//fLogger
	/* go func() {
		fLogger.Write(logEcho)
	}() */

	//MgoClient
	go func() {
		client := mgo.MongoClient().Copy()
		defer client.Close()

		if err := client.DB("document").C(lg.Collection).Insert(&lg); err != nil {
			fmt.Printf("\n err Logs time:%s, %s\n", time.Now(), err)
		} else {
			//fmt.Printf("\n not err, time:%s\n", time.Now())
		}
	}()

	return len(logEcho), nil
}

//Zaplog Log
func (lg *Zaplog) Write(logByte []byte) (n int, err error) {

	if err := json.Unmarshal([]byte(logByte), &lg); err != nil {
		fmt.Println("Unmarshal err =>", err)
		//return err
		//panic(err)
	}

	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}

		//log.Printf("response: %q", lg)
		return len(logByte), nil
		//return err
	}
	//fLogger
	/* go func() {
		fLogger.Write(logByte)
	}() */

	//MgoClient
	go func() {
		client := mgo.MongoClient().Copy()
		defer client.Close()

		if err := client.DB("document").C(lg.Collection).Insert(&lg); err != nil {
			fmt.Printf("\n err Logs time:%s, %s\n", time.Now(), err)
		}
	}()

	return len(logByte), nil
}

// ZapLogger is an example of echo middleware that logs requests using logger "zap"
func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(
		middleware.BodyDumpConfig{
			Handler: func(c echo.Context, reqBody, resBody []byte) {

				start := time.Now()

				req := c.Request()
				res := c.Response()

				reqB := `""`

				reqForm := req.Form
				jsonString, err := json.Marshal(reqForm)
				if err != nil {
					fmt.Println("err jsonString  => ", err)
				}

				if string(jsonString) != "null" {
					reqB = string(jsonString)
				} else if len(reqBody) > 0 {
					reqB = string(reqBody)
				}

				id := req.Header.Get(echo.HeaderXRequestID)
				if id == "" {
					id = res.Header().Get(echo.HeaderXRequestID)
				}

				jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)

				jsonData := make(map[string]interface{})
				if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
					fmt.Println("err jsonData  => ", err)
				}

				fields := []zapcore.Field{
					zap.Int("status", res.Status),
					zap.String("latency", time.Since(start).String()),
					zap.String("id", id),
					zap.String("method", req.Method),
					zap.String("uri", req.RequestURI),
					zap.String("host", req.Host),
					zap.String("remote_ip", c.RealIP()),
					zap.String("prefix", "API Log"),
					zap.String("message", "{}"),
					zap.Any("data", jsonData),
				}

				n := res.Status
				switch {
				case n >= 500:
					log.Error("Server error", fields...)
				case n >= 400:
					log.Warn("Client error", fields...)
				case n >= 300:
					log.Info("Redirection", fields...)
				default:
					log.Info("Success", fields...)
				}

				//return nil
			},
		})
}

// TimeEncoder return time encode
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z"))
}

//InitialLogs is init logs
func InitialLogs(e *echo.Echo) {

	fmt.Println("InitialLogs..")

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&Zaplog{Collection: "logger"}),
			zapcore.AddSync(&lumberjack.Logger{
				Filename: filepath.Join("./logs", fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), ".log")), //	Format YYYY-MM-DD
				MaxSize:  100,                                                                                   // megabytes
				MaxAge:   28,                                                                                    //	days
				Compress: true,                                                                                  // disabled by default
			}),
		),
		zap.InfoLevel,
	)

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	//log.SetFlags(0)

	zaplogger := zap.New(core, zap.AddCaller())

	e.Use(ZapLogger(zaplogger))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"level":"info", "time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\r\n",
		Output: &Logs{Collection: "logs"},
	}))
}