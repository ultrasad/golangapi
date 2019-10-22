package middlewares

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	//"context"
	//"golangapi/db/mongo"
	"golangapi/db/mgo"
	//"github.com/globalsign/mgo/bson"
	"log"
	"strings"
	"time"

	//"bytes"
	//"encoding/base64"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	collection string
	err        error
	fLogger    *lumberjack.Logger
	//errLog     *log.Logger
	//errLog *lumberjack.Logger
)

type (
	//Logger struct logger from go log
	Logger struct {
		Time       time.Time `bson:"time" json:"time"`
		Lv         string    `bson:"level" json:"level"`
		Prefix     string    `bson:"prefix" json:"prefix"`
		Message    string    `bson:"-" json:"message"`
		Data       CtxLogger `bson:"data" json:"data"`
		Collection string    `bson:"-"`
	}

	//ReqLogger request logger
	/* ReqLogger struct {
		Fake bool  `json:"fake" bson:"fake"`
		Field1 string  `json:"field1" bson:"field1"`
		Status int  `json:"status" bson:"status"`
		Topic string  `json:"topic" bson:"topic"`
	} */

	//CtxRrequest struct logger request
	CtxRrequest struct {
		Fake   bool   `bson:"fake" json:"fake"`
		Field1 string `bson:"field1" json:"field1"`
		Status int    `bson:"status" json:"status"`
		Topic  string `bson:"topic" json:"topic"`
	}

	//CtxLogger struct logger req,res
	CtxLogger struct {
		ID  string      `bson:"id" json:"id"`
		Req interface{} `bson:"request" json:"req"`
		Res interface{} `bson:"response" json:"res"`
		//Req       ReqLogger `bson:"req" json:"req"`
		//Res interface{} `bson:"response" json:"res"`
		//Req json.RawMessage
		//Res json.RawMessage
		//Req map[string]interface{} `bson:"request" json:"req"`
		//Res map[string]interface{} `bson:"response" json:"res"`
		//Req map[string]interface{} `bson:"request" json:"req"`
		//Req CtxRrequest            `bson:"request" json:"req"`
		//Res map[string]interface{} `bson:"response" json:"res"`
	}

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

	//Logrus struct log from Logrus
	Logrus struct {
		// Time       time.Time   `json:"time" bson:"time"`
		// Animal     string      `json:"animal" bson:"animal"`
		// Data       ctxLogger   `bson:"data" json:"data"`
		// ID         string      `json:"id" bson:"id"`
		// Req        interface{} `json:"req" bson:"request"`
		// Res        interface{} `json:"res" bson:"response"`
		// Message    string      `bson:"-" json:"message"`
		// Collection string      `bson:"-"`

		Time       time.Time `bson:"time" json:"time"`
		Lv         string    `bson:"level" json:"level"`
		Prefix     string    `bson:"prefix" json:"prefix"`
		Message    string    `bson:"-" json:"message"`
		Msg        string    `bson:"msg" json:"msg"`
		Data       CtxLogger `bson:"data" json:"data"`
		Collection string    `bson:"-"`
	}

	//Zaplog struct log from Zap logger
	Zaplog struct {
		// Time       time.Time   `json:"time" bson:"time"`
		// Animal     string      `json:"animal" bson:"animal"`
		// Data       ctxLogger   `bson:"data" json:"data"`
		// ID         string      `json:"id" bson:"id"`
		// Req        interface{} `json:"req" bson:"request"`
		// Res        interface{} `json:"res" bson:"response"`
		// Message    string      `bson:"-" json:"message"`
		// Collection string      `bson:"-"`

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

//InitLog init logger
func InitLog() {

	//some time shutdown database, you will need this.
	//year, month, day := time.Now().Date()
	/*fLogger = &lumberjack.Logger{
		Filename: filepath.Join("./logs", strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)+".log"),
		MaxSize:  650,  // megabytes
		MaxAge:   15,   //days
		Compress: true, // disabled by default
	}*/

	// Log as JSON instead of the default ASCII formatter.
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	//mWriter := io.MultiWriter(os.Stderr, &middlewares.Logrus{Collection: "logger"})
	//logrus.SetOutput(mWriter)
	//logrus.SetOutput(&middlewares.Logrus{Collection: "logger"})

	//logrus.SetLevel(logrus.WarnLevel)
	//logrus.SetLevel(logrus.InfoLevel)

	// Caller func report
	//logrus.SetReportCaller(true)

	//year, month, day := time.Now().Date()
	fLogger = &lumberjack.Logger{
		//Filename: filepath.Join("./logs", strconv.Itoa(year)+"-"+strconv.Itoa(int(month))+"-"+strconv.Itoa(day)+".log"),
		//Dir:        "./logs",
		//NameFormat: time.RFC822 + ".log",
		Filename: filepath.Join("./logs", fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), ".log")), //	Format YYYY-MM-DD
		MaxSize:  100,                                                                                   // megabytes
		MaxAge:   15,                                                                                    //	days
		Compress: true,                                                                                  // disabled by default
	}

	fmt.Println("init logs..." + fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), ".log"))

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	//log.SetFlags(0)

}

func (lg *Logger) Write(logByte []byte) (n int, err error) {

	//fmt.Println("log", lg)
	// fmt.Println("logByte", logByte)

	// //var dataX interface{}
	// var dataX interface{}
	// if err = json.Unmarshal(logByte, dataX); err != nil {
	// 	log.Printf("err dataX: %q", err)
	// } else {
	// 	log.Printf("dataX: %q", dataX)
	// }

	err = json.Unmarshal(logByte, &lg)
	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		//log.Printf("response: %q", lg)
		return len(logByte), nil
		//return err
	}

	/*
		if err != nil {
			fmt.Println("\n err Logger, json Unmarshal >>>", err)
			return
		}
	*/

	//fmt.Println("lg.Message => ", lg.Message)
	//fmt.Println("lg.Data => ", lg.Data)

	if err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data); err != nil {
		return
	}

	/* MongoClient */
	/*
		go func() {
			client := mongo.ClientManager()
			// create a new context with a 10 second timeout
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			insertResult, err := client.Database("document").Collection(lg.Collection).InsertOne(ctx, &lg)
			if err != nil {
				fmt.Printf("\n err time:%s,%s\n", time.Now(), lg.Message)
			}
			fmt.Println("Inserted a Logger: ", insertResult.InsertedID)
		}()
	*/

	//MgoClient
	go func() {
		client := mgo.MongoClient().Copy()
		defer client.Close()

		if err := client.DB("document").C(lg.Collection).Insert(&lg); err != nil {
			fmt.Printf("\n err time:%s,%s\n", time.Now(), lg.Message)
		} else {
			//fmt.Printf("\n not err, time:%s\n", time.Now())
		}
	}()

	return len(logByte), nil
}

/*func LoggerLumberjack() *lumberjack.Logger {
	return fLogger
}*/

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
	// MongoClinet
	/*
		go func() {
			client := mongo.ClientManager()
			// create a new context with a 10 second timeout
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			insertResult, err := client.Database("document").Collection(lg.Collection).InsertOne(ctx, &lg)
			if err != nil {
				fmt.Printf("\n err time:%s,%s\n", time.Now(), err)
			}
			fmt.Println("Inserted a Log: ", insertResult.InsertedID)
		}()
	*/

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

//Logrus Log
func (lg *Logrus) Write(logByte []byte) (n int, err error) {
	//func (lg *Logrus) Write(logByt interface{}) (n int, err error) {

	//fmt.Println("log Logrus byte => ", logByte)
	// var f map[string]interface{}
	// if err := json.Unmarshal(logByte, &f); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("log all => ", &f)

	//err = json.Unmarshal(logByte, &f)
	err = json.Unmarshal(logByte, &lg)
	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}

		//log.Printf("response: %q", lg)
		return len(logByte), nil
		//return err
	}

	//fmt.Println("Logrus lg => ", &lg)
	//fmt.Println("lg Interface Data => ", &lg)
	//fmt.Println("lg.Message => ", lg.Message)
	//fmt.Println("lg.Data => ", lg.Data)

	//var jsonData map[string]interface{}

	//fmt.Println("lg => ", &lg.Data)
	//err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)
	err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)
	if err != nil {
		fmt.Println("\n err json decode >>>", err)
		return
	}

	//fmt.Println("lg => ", &lg.Data)
	//fmt.Printf("Message: %s Data: %s", lg.Message, lg.Data)
	//fmt.Printf("Response: %s", lg.Data)
	//fmt.Printf("lg.Collection: %s", lg.Collection)

	//fLogger
	/*go func() {
		fLogger.Write(logByte)
	}()*/

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

/* func (s Zaplog) Len() int {
    return len(s)
}

func (s Zaplog) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s Zaplog) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
} */

// Zaplogs implements sort.Interface for []Zaplog based on
// the Lv field.
/* type Zaplogs struct {
	Zaplogs []Zaplog
} */

/*
type reqByName []Zaplog

// Ensure it satisfies sort.Interface
func (d reqByName) Len() int           { return len(d) }
func (d reqByName) Less(i, j int) bool { return d[i].Message < d[j].Message }
func (d reqByName) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

type reqByField []CtxLogger

// Ensure it satisfies sort.Interface
func (d reqByField) Len() int           { return len(d) }
func (d reqByField) Less(i, j int) bool { return d[i].ID < d[j].ID }
func (d reqByField) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
*/

//Len ByAge
//func (a ByAge) Len() int { return len(a) }
//Swap ByAge
//func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
//Less ByAge
//func (a ByAge) Less(i, j int) bool { return a[i].Time < a[j].Time }

//ToSlice Zaplog
/* func ToSlice(m map[string]Zaplog) []Zaplog {
    cities := make([]Zaplog, 0, len(m))
    for k, v := range m {
        v.Data = k
        cities = append(cities, v)
    }
    return cities
} */

//UnmarshalJSON logger CtxLogger
/* func (cm *CtxLogger) UnmarshalJSON(bs []byte) error {
	// Unquote the source string so we can unmarshal it.
	//unquoted, err := strconv.Unquote(string(bs))
	//if err != nil {
	//	fmt.Println("unquoted =>", err)
	//	return err
	//}

	//fmt.Println("unquoted =>", unquoted)

	// Create an aliased type so we can use the default unmarshaler.
	type CustomMeta2 CtxLogger
	var cm2 CustomMeta2

	// Unmarshal the unquoted string and assign to the original object.
	//if err := json.Unmarshal([]byte(unquoted), &cm2); err != nil {
	if err := json.Unmarshal([]byte(bs), &cm2); err != nil {
		fmt.Println("Unmarshal =>", err)
		return err
	}

	fmt.Println("cm2 => ", &cm2)

	jsonBytesReq, _ := json.MarshalIndent(&cm2.Req, "", "  ")
	//fmt.Println(countries)
	fmt.Println("jsonBytesReq => ", string(jsonBytesReq))

	json.Unmarshal(jsonBytesReq, &cm2.Req)

	jsonBytesRes, _ := json.MarshalIndent(&cm2.Res, "", "  ")
	//fmt.Println(countries)
	fmt.Println("jsonBytesRes => ", string(jsonBytesRes))

	json.Unmarshal(jsonBytesRes, &cm2.Res)

	*cm = CtxLogger(cm2)

	fmt.Println("cm => ", &cm)

	//jsonBytes, _ := json.MarshalIndent(cm.Req, "", "  ")
	//fmt.Println(countries)
	//fmt.Println("jsonBytes => ", string(jsonBytes))

	return nil
} */

//Zaplog Log
func (lg *Zaplog) Write(logByte []byte) (n int, err error) {
	//func (lg *Logrus) Write(logByt interface{}) (n int, err error) {

	//fmt.Println("log Logrus byte => ", logByte)
	// var f map[string]interface{}
	// if err := json.Unmarshal(logByte, &f); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("log all => ", &f)

	/* jsonData := make(map[string]interface{})
	//var jsonData interface{}
	err = json.Unmarshal([]byte(string(logByte[:])), &jsonData)

	fmt.Println("Zaplog jsonData => ", &jsonData) */
	//fmt.Println("Zaplog jsonData str => ", jsonData["str"])
	//mFriends := mResult["friends"].(map[int]map[string]interface{})
	//jData := jsonData["str"].(map[string]interface{})

	//changes := bson.M(jsonData)

	/* jsonData := make(map[string]interface{})
	err = json.Unmarshal([]byte(logByte), &jsonData) */

	//sort.Sort(ByAge(&people))
	//fmt.Println("jsonData => ", &people)

	//jsonStr := make(map[string]interface{})
	//jsonStr["data"] = jsonData["str"]
	//json.Marshal(jsonStr)

	/* 	var f interface{}
	err = json.Unmarshal(logByte, &f)

	if err != nil {
		panic("OMG!!")
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		fmt.Printf("%s => %s", k, v)
	}

	fmt.Println("res m => ", &m) */

	//r := bytes.NewReader(logByte)
	//dec := base64.NewDecoder(base64.StdEncoding, r)

	/* var obj Zaplog
	if err := json.Unmarshal(logByte, &obj); err != nil {
		panic(err)
	} */

	/* var people []*Zaplog

	   likes := make(map[string][]*Zaplog)
	   for _, p := range people {
	       for _, l := range p.Data {
	           likes[l] = append(likes[l], p)
	       }
	   } */

	//fmt.Println("obj zaplog => ", obj)

	//err = json.Unmarshal(logByte, &f)
	//err = json.Unmarshal(logByte, &lg)
	//err = json.Unmarshal([]byte(logByte), &lg)
	//sort.Sort(&lg.Data)

	/* var zaplogsX []Zaplogs
		err = json.Unmarshal([]byte(logByte), &zaplogsX)
		sort.Sort(zaplogsX)

	    for _, d := range zaplogsX {
	        fmt.Println("d => ", d)
		} */

	//jsonData := make(map[string]interface{})
	//var jsonData interface{}
	//var jsonData []Zaplog
	//var newlog []Zaplog
	//var newlg []Zaplog

	//jsonData := make(map[string]interface{})
	//var jsonData []Zaplog

	/* jsonData := Zaplog{}
	err = json.Unmarshal([]byte(logByte), &jsonData)

	vehicles := make(Zaplogs, 0, len(jsonData))
	fmt.Println("Zaplog jsonData vehicles => ", vehicles)
	for _, c := range jsonData {
		vehicles = append(vehicles, Zaplog(c))
		//fmt.Println("Zaplog c => ", c)
	}
	sort.Sort(vehicles) */

	//var vData interface{}
	//json.Unmarshal(logByte, &vData)
	//data := vData.(map[string]interface{})

	// Unquote the source string so we can unmarshal it.
	/* unquoted, err := strconv.Unquote(string(logByte))
	if err != nil {
		//return err
		panic(err)
	} */

	//fmt.Println("unquoted => ", unquoted)

	//jsonData := make(map[string]interface{})
	//json.Unmarshal([]byte(logByte), &jsonData)

	//err = json.Unmarshal([]byte(logByte), &lg)
	//fmt.Println("Zaplog jsonData lg => ", &lg.Data)

	//fmt.Println("vData => ", vData.Data.Req)

	/* err = json.Unmarshal([]byte(logByte), &lg)

	jsonData := make(map[string]interface{})
	err = json.Unmarshal([]byte(logByte), &jsonData) */

	//json.Unmarshal([]byte(logByte), &lg)

	//var doc Zaplog
	//jsonData := Zaplog{}
	//json.Unmarshal([]byte(logByte), &jsonData)
	/* err := json.Unmarshal([]byte(logByte), &doc)
	if err != nil {
		panic(err)
	} */

	/*sort.Slice(jsonData.Data.Req, func(i, j int) bool {
		//p1 := jsonData.Data.Req[i]
		//p2 := jsonData.Data.Req[j]
		//return p1 < p2
		fmt.Printf("i:%v, j:%v", i, j)
		return i < j
	})*/

	/* err = json.NewDecoder(strings.NewReader(string(logByte))).Decode(&lg)
	if err != nil {
		fmt.Println("\n err json decode >>>", err)
		return
	} */

	//jsonData := make(map[string]interface{})
	//jsonData := Zaplog{}
	//type CustomMeta2 Zaplog
	//var cm2 CustomMeta2

	if err := json.Unmarshal([]byte(logByte), &lg); err != nil {
		fmt.Println("Unmarshal err =>", err)
		//return err
		//panic(err)
	}

	//fmt.Println("&lg => ", &lg)

	//jsonBytesReq, _ := json.MarshalIndent(&cm2.Data.Req, "", "  ")
	//fmt.Println(countries)
	//fmt.Println("jsonBytesReq => ", string(jsonBytesReq))

	//err = json.Unmarshal([]byte(logByte), &cm2)
	//fmt.Println("cm2 => ", cm2)

	//jsonBytes, _ := json.MarshalIndent(&cm2.Data.Req, "", "  ")
	//fmt.Println("jsonBytes => ", jsonBytes)

	//err = json.Unmarshal(jsonBytesReq, &cm2.Data.Req)

	//fmt.Println("cm2 after => ", &cm2)

	//Sort key OK
	/* fmt.Println("jsonData => ", lg)
	var keys []string
	//for _, currMap := range lg.Data.Req {
	for k, v := range lg.Data.Req {
		lg.Data.Req[k] = v
		keys = append(keys, k)
	}
	//}
	fmt.Println("keys before => ", keys)

	sort.Strings(keys)

	fmt.Println("keys after => ", keys) */

	/* sort.Slice(lg.Data.Req, func(i, j int) bool {
		//p1 := doc.MetaInfos[i].Custom.PartNum
		//p2 := doc.MetaInfos[j].Custom.PartNum
		//return p1 < p2
		//return lg.Data.Req[i] < lg.Data.Req[j]
	}) */

	/* for _, m := range doc.MetaInfos {
		fmt.Println(m.Custom.PartNum)
	} */

	/* keys := make([]string, 0, len(jsonData))
	for k := range jsonData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, " => ", jsonData[k])
	}

	//json.Unmarshal([]byte(logByte), &lg)
	//sort.Sort(reqByName(lg))

	vData := reqByName{}

	fmt.Println("vData => ", vData)

	sort.Sort(vData)

	fmt.Println("vData sort => ", vData)

	sort.Slice(vData, func(p, q int) bool {
		return vData[p].Data.ID < vData[q].Data.ID
		//return p < q
		//fmt.Printf("%v => %v", p, q)
		//return true
	})

	fmt.Println("vData slice => ", vData) */

	/* sort.Slice(vData.Data.Req, func(p, q int) bool {
		//return p < q
		//return p < q
		fmt.Printf("%v => %v", p, q)
		return true
	}) */

	/* sort.Slice(vData.Data.Req, func(p, q int) bool {
		return p < q
		//return p < q
	})

	fmt.Println("vData sorted => ", vData) */

	// Sorting Author by their name
	// Using Slice() function
	/* sort.Slice(&lg, func(p, q int) bool {
		return &lg[p]. < Author[q].a_name })
		//return p < q
	})

	fmt.Println("Sort Author according to their names:")
	fmt.Println("&lg.Data.Req => ", &lg) */

	/* res1 := sort.SliceIsSorted(&lg.Data.Req, func(p, q int) bool {
		//return &lg.Data.Req[p] < &lg.Data.Req[q]
		//fmt.Printf("%v => %v", p, q)
		return p < q
	}) */

	//fmt.Println("sort.IsSorted => ", res1)

	//fmt.Println("sort.IsSorted => ", sort.StringSlice(&lg.Data.Req))

	if err != nil {
		log.Printf("error decoding response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}

		//log.Printf("response: %q", lg)
		return len(logByte), nil
		//return err
	}

	/* fmt.Println("Zaplog lg => ", &lg)
	fmt.Println("lg Interface Data => ", &lg)
	fmt.Println("lg.Message => ", lg.Message)
	fmt.Println("lg.Data => ", lg.Data) */

	//var jsonData map[string]interface{}

	//fmt.Println("json data to mongo before >>>", &lg)

	//fmt.Println("lg => ", &lg.Data)
	//err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)
	/* err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)
	if err != nil {
		fmt.Println("\n err json decode >>>", err)
		return
	} */

	/* err = json.NewDecoder(strings.NewReader(lg.Message)).Decode(&lg.Data)
	if err != nil {
		fmt.Println("\n err json decode >>>", err)
		return
	} */

	//fmt.Println("json data to mongo after => ", &lg)

	/* fmt.Println("lg => ", &lg.Data)
	fmt.Printf("Message: %s Data: %s", lg.Message, lg.Data)
	fmt.Printf("Response: %s", lg.Data)
	fmt.Printf("lg.Collection: %s", lg.Collection) */

	//fLogger
	go func() {
		fLogger.Write(logByte)
	}()

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
