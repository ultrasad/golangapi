package controllers

import (
	"context"
	"strings"
	"bytes"
	"encoding/json"
	"log"
	"github.com/labstack/echo"
	"net/http"
	"fmt"

	"golangapi/models"

	elastics "golangapi/db/elastics"
)

//Search return search result
func Search(c echo.Context) (err error){

	//ES Client
	es := elastics.ESClient()

	info, _ := es.Info()
	fmt.Println("ES Client Info",info)
	
	// 3. Search for the indexed documents
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				//"customer_refno": "CS0900528",
				"customer_mobile": "66851004508",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("customers"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithSize(10),
		es.Search.WithFrom(0),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	//fmt.Println("res.Body => ", res.Body)

	//Other
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	//var response map[string]interface{}
	//var customer models.Customer

	r := make(map[string]interface{})
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	var customers []*models.Customer

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"].(map[string]interface{})["firstname"])
		//log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])

		customers = append(customers, &models.Customer{
			Firstname: hit.(map[string]interface{})["_source"].(map[string]interface{})["firstname"].(string),
			Lastname: hit.(map[string]interface{})["_source"].(map[string]interface{})["lastname"].(string),
		})
	}

	log.Println(strings.Repeat("~", 37)) //print ~

	fmt.Println("customers => ", customers)

	//var result map[string]interface{}
	//result["_source"] = r["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"]


	//var customer models.Customer
	
	//return c.JSON(http.StatusOK, r["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"])
	//return c.JSON(http.StatusOK, r["hits"].(map[string]interface{})["hits"].([]interface{})[0])
	//return c.JSON(http.StatusOK, response["hits"].(map[string]interface{})["hits"].([]interface{}))
	//return c.JSON(http.StatusOK, response["hits"])
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"total": int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		"took": int(r["took"].(float64)),
		"customers": r["hits"].(map[string]interface{})["hits"].([]interface{}),
	})
  	//return err
}