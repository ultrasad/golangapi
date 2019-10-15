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

				/*
					if err := json.Unmarshal([]byte(reqForm), &body); err != nil {
						//return nil, err
						fmt.Println("err Unmarshal => ", err)
					}

					fmt.Println("body => ", body)
				*/

				start := time.Now()

				// err := next(c)
				// if err != nil {
				// 	c.Error(err)
				// }

				//fmt.Printf("Request Body: %v\n", string(reqBody))
				//fmt.Printf("Response Body: %v\n", string(resBody))

				/*reqB := `""`
				resB := string(resBody)
				if len(reqBody) > 0 {
					reqB = string(reqBody)
				}

				req := c.Request()
				res := c.Response()

				id := req.Header.Get(echo.HeaderXRequestID)
				if id == "" {
					id = res.Header().Get(echo.HeaderXRequestID)
				}*/

				//fmt.Println("reqBody  => ", reqBody)

				/* reqB := `""`
				if len(reqBody) > 0 {
					reqB = string(reqBody)
				} */

				req := c.Request()
				res := c.Response()

				reqB := `""`
				//reqForm := c.Request().Form
				//reqForm := c.Request().Form
				reqForm := req.Form
				jsonString, err := json.Marshal(reqForm)
				if err != nil {
					fmt.Println("err jsonString  => ", jsonString)
				}

				//fmt.Println("jsonString  => ", jsonString)
				//fmt.Printf("length => %v", len(jsonString))

				if string(jsonString) != "null" {
					//fmt.Println("jsonString not nil => ", string(jsonString))
					reqB = string(jsonString)
				} else if len(reqBody) > 0 {
					reqB = string(reqBody)
				}

				//fmt.Println("reqB  => ", reqB)

				id := req.Header.Get(echo.HeaderXRequestID)
				if id == "" {
					id = res.Header().Get(echo.HeaderXRequestID)
				}

				// m := echo.Map{}
				// if err := c.Bind(&m); err != nil {
				// 	//return err
				// 	//panic(err)
				// 	fmt.Println("m err => ", err)
				// }

				// fmt.Println("m => ", m)

				//fmt.Println("reqBody => ", reqBody)
				//fmt.Println("resBody => ", resBody)

				/*
					jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
					//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), req.Body, resBody)
					fmt.Println("jsonStr before => ", jsonStr)

					var p2 interface{}
					json.Unmarshal([]byte(jsonStr), &p2)
					jsonData := p2.(map[string]interface{})
					fmt.Println("jsonData after => ", jsonData)
				*/

				/*
					jsonData := echo.Map{}
					if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
						panic(err)
					}
				*/

				// var dataX interface{}
				// if err := json.NewDecoder(c.Request().Body).Decode(&dataX); err != nil {
				// 	panic(err)
				// } else {
				// 	fmt.Println("dataX after => ", dataX)
				// }

				// jsonMap := make(map[string]interface{})
				// //err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
				// err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonMap)
				// if err != nil {
				// 	//return err
				// 	panic(err)
				// } else {
				// 	fmt.Println("jsonMap after => ", jsonMap)
				// }

				// err = json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonData)
				// if err != nil {
				// 	fmt.Println("\n err json decode >>>", err)
				// 	return
				// }

				// resX := jsonData["res"].(map[string]interface{})

				// for key, value := range resX {
				// 	// Each value is an interface{} type, that is type asserted as a string
				// 	//fmt.Println(key, value.(string))
				// 	fmt.Println(key, value)
				// }

				//fmt.Println("jsonStr after => ", jsonData)

				jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resBody)
				//jsonStr := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), reqB, resB)

				//fmt.Println("jsonStr before => ", jsonStr)

				//jsonData := make(map[string]interface{})
				//var jsonData map[string]interface{}
				jsonData := make(map[string]interface{})
				//var jsonData interface{}
				if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
					//panic(err)
					//var jsonData2 interface{}

					//var jsonMap map[string]interface{} = map[string]interface{}{}
					/* var jsonMap map[string]interface{}
					if err := c.Bind(&jsonMap); err != nil {
						//return err
						fmt.Println("err jsonMap => ", jsonMap)
					} */

					/* r := strings.NewReader(reqB)
					b, err := ioutil.ReadAll(r)
					if err != nil {
						//log.Fatal(err)
						fmt.Println("err jsonMap => ", r)
					}

					fmt.Printf("jsonMap => %s", b) */

					//fmt.Println("jsonMap => ", jsonMap)

					//fmt.Println("panic jsonStr => ", jsonStr)

					/*
						var reqForm = c.Request().Form
						fmt.Println("panic jsonData, reqForm => ", reqForm)

						jsonString, err := json.Marshal(reqForm)
						if err != nil {
							fmt.Println("err jsonString  => ", jsonString)
						}

						//mBody := make(map[string]interface{})
						//for key, value := range reqForm {
						//	fmt.Printf("%s = %s \r\n", key, value[0])
						//	jsonData[key] = value[0]
						//}

						fmt.Println("jsonString => ", string(jsonString))

						jsonStrNew := fmt.Sprintf(`{"id":"%s","req":%s,"res":%s}`, c.Response().Header().Get(echo.HeaderXRequestID), string(jsonString), resBody)
						if err := json.Unmarshal([]byte(jsonStrNew), &jsonData); err != nil {
							fmt.Println("err jsonStrNew => ", string(jsonString))
						}

						fmt.Println("jsonData => ", jsonData)
					*/

				}

				//fmt.Println("jsonStr after => ", jsonData)

				/* jsonData := make(map[string]interface{})
				jsonData["id"] = c.Response().Header().Get(echo.HeaderXRequestID)
				jsonData["req"] = reqB
				jsonData["res"] = resBody

				fmt.Println("jsonStr after => ", jsonData) */

				/* jsonData := make(map[string]interface{})
				err := json.Unmarshal([]byte(jsonStr), &jsonData)

				if err != nil {
					panic(err)
				} */

				/* var keys []string
				for key, value := range jsonData {
					keys = append(keys, key)
					fmt.Println("index : ", key, " value : ", value)
				}

				sort.Strings(keys)

				// To perform the opertion you want
				for _, k := range keys {
					fmt.Println("Key:", k, "Value:", jsonData[k])
				} */

				//jsonData := make(map[string]interface{})
				/*var jsonData interface{}

				err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonData)
				if err != nil {
					fmt.Println("\n err json decode >>>", err)
					return
				}*/

				//fmt.Println("jsonData after => ", jsonData)

				// m := echo.Map{}
				// if err := c.Bind(&m); err != nil {
				// 	//return err
				// 	panic(err)
				// }
				// fmt.Printf("map => %+v", m)

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
