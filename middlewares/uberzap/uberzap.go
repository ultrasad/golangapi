package uberzap

import (
	"os"

	"go.uber.org/zap"
)

// LoggerPathFormatter return string format path
type LoggerPathFormatter func() string

//NewLogger stamp logs and log output
func NewLogger(path string, callback LoggerPathFormatter) (*zap.Logger, error) {
	os.Mkdir(path, os.ModePerm)

	config := zap.NewProductionConfig()
	/*config.OutputPaths = []string {
		callback(),
	}*/

	config.OutputPaths = []string{
		callback(),
		"stdout",
	}

	config.EncoderConfig.LevelKey = "level"
	//config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.TimeKey = ""
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.MessageKey = "message"

	return config.Build()
}
