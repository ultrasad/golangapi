package handler

import (
	"fmt"
	"golangapi/middlewares"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TimeEncoder return time encode
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(t.Format("2006-01-02T15:04:05Z07:00"))
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05Z"))
}

//InitialLogsOld is init logs
func InitialLogsOld(e *echo.Echo) {

	//e := echo.New()

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
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&middlewares.Logrus{Collection: "logger"})), //old ok
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&middlewares.Zaplog{Collection: "logger"})),
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(&middlewares.Logrus{Collection: "logger"})),
		//zap.DebugLevel,
		zap.InfoLevel,
	)

	zaplogger := zap.New(core, zap.AddCaller())

	e.Use(middlewares.ZapLogger(zaplogger))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"level":"info", "time":"${time_rfc3339}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\r\n",
		//CustomTimeFormat: "2006-01-02T15:04:05Z",
		Output: &middlewares.Logs{Collection: "logs"},
		//Output: os.Stdout,
		//Output: echoLog,
	}))

	fmt.Println("init logs...")

}
