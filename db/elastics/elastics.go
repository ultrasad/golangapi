package elastics

import (
	"github.com/spf13/viper"
	"net/http"
	"time"
	"net"
	"crypto/tls"
	"log"
	"fmt"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

var (
	esClient *elasticsearch.Client
	err   error
)

// ConnectES return Mongo Connection
func ConnectES() *elasticsearch.Client {
	ESHost := viper.GetString("elastics.url")
	cfg := elasticsearch.Config {
		Addresses: []string{
		  //"http://localhost:9200",
		  //"http://localhost:9201",
		  ESHost,
		},
		//Username: "foo",
		//Password: "bar",
		Transport: &http.Transport{
		  MaxIdleConnsPerHost:   10,
		  ResponseHeaderTimeout: time.Second,
		  DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		  TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS11,
		  },
		},
	}
	
	esClient, err = elasticsearch.NewClient(cfg)
	
	if(err != nil){
		log.Fatalf("Error creating the client: %s\n", err)
	}
	
	info, _ := esClient.Info()
	fmt.Println("ES Info",info)

	return esClient
}

// ESClient return MongoDB Session
func ESClient() *elasticsearch.Client {
	fmt.Println("Call ES Client... ")
	return esClient
}