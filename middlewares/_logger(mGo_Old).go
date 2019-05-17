package middlewares

import (
	"encoding/json"
	"fmt"
	"golangapi/db/mongo"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	collection string
	err        error
	fLogger    *lumberjack.Logger
	//errLog     *log.Logger
)

type (
	//Logger struct logger from go log
	Logger struct {
		Time       time.Time `bson:"time" json:"time"`
		Lv         string    `bson:"level" json:"level"`
		Prefix     string    `bson:"prefix" json:"prefix"`
		Message    string    `bson:"-" json:"message"`
		Data       ctxLogger `bson:"data" json:"data"`
		Collection string    `bson:"-"`
	}

	ctxLogger struct {
		ID  string      `json:"id" bson:"id"`
		Req interface{} `json:"req" bson:"request"`
		Res interface{} `json:"res" bson:"response"`
	}

	// Logs struct log from echo
	Logs struct {
		ID           string    `json:"id" bson:"id"`
		Time         time.Time `json:"time" json:"time"`
		RemoteIP     string    `json:"remote_ip" bson:"remote_ip"`
		Host         string    `json:"host" bson:"host"`
		Method       string    `json:"method" bson:"method"`
		URI          string    `json:"uri" bson:"uri"`
		Status       int       `json:"status" bson:"status"`
		Latency      int       `json:"latency" bson:"latency"`
		LatencyHuman string    `json:"latency_human" bson:"latency_human"`
		BytesIn      int       `json:"bytes_in" bson:"bytes_in"`
		BytesOut     int       `json:"bytes_out" bson:"bytes_out"`
		Collection   string    `bson:"-"`
	}
)

//Init log
func init() {

	//some time shutdown database, you will need this.
	year, month, day := time.Now().Date()
	fLogger = &lumberjack.Logger{
		Filename: filepath.Join("./logs", strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)+".log"),
		MaxSize:  650,  // megabytes
		MaxAge:   15,   //days
		Compress: true, // disabled by default
	}

	fmt.Println("init logs...")
}

func (lg *Logger) Write(logByte []byte) (n int, err error) {

	err = json.Unmarshal(logByte, &lg)
	if err != nil {
		fmt.Println("\n err Logger, json Unmarshal >>>", err)
		return
	}

	err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)

	if err != nil {
		//fmt.Println("\n err json decode >>>", err)
		return
	}

	go func() {
		conn := mongo.MgoManager().Copy()
		defer conn.Close()

		if err := conn.DB("document").C(lg.Collection).Insert(&lg); err != nil {
			fmt.Printf("\n err time:%s,%s\n", time.Now(), lg.Message)
		} else {
			//fmt.Printf("\n not err, time:%s\n", time.Now())
		}
	}()
	return len(logByte), nil
}

//echo Logs
// 2019-05-14, comment fix test
func (lg *Logs) Write(logEcho []byte) (n int, err error) {

	err = json.Unmarshal(logEcho, &lg)
	if err != nil {
		fmt.Println("\n err Logs, json Unmarshal >>>", err)
		return
	}

	//fLogger
	go func() {
		fLogger.Write(logEcho)
	}()

	//fmt.Printf("\n &lg Logs: %#v\n", &lg)

	go func() {
		conn := mongo.MgoManager().Copy()
		defer conn.Close()

		if err := conn.DB("document").C(lg.Collection).Insert(&lg); err != nil {
			fmt.Printf("\n err Logs time:%s, %s\n", time.Now(), err)
		} else {
			//fmt.Printf("\n not err, time:%s\n", time.Now())
		}
	}()

	return len(logEcho), nil
}
