package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is an example of echo middleware that logs requests using logger "zap"
/*
func ZapLogger(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			fields := []zapcore.Field{
				zap.Int("status", res.Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("id", id),
				zap.String("method", req.Method),
				zap.String("uri", req.RequestURI),
				zap.String("host", req.Host),
				zap.String("remote_ip", c.RealIP()),
				//zap.String("request", req),
				//zap.String("request", res),
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

			return nil
		}
	}
}
*/

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

				fmt.Println("jsonStr  => ", jsonStr)

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
					//zap.String("str", jsonStr),
					zap.Any("data", jsonData),
					//zap.Binary("data", []byte(jsonStr)),
					//zap.String()
					//zap.ByteString("data", resBody),
					//zap.Any("data", jsonData),
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
