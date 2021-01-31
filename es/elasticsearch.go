package es

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var esClient *elasticsearch.Client = nil
var r map[string]interface{}

func Init() {
	if esClient == nil{
		config := elasticsearch.Config{
			Addresses: []string{
				"http://localhost:9200",
			},
		}
		var err error = nil
		esClient, err = elasticsearch.NewClient(config)
		if err != nil {
			log.Fatalf("Error creating the client: %s", err)
		}

		// Get cluster info
		res, err := esClient.Info()
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		// Check response status
		if res.IsError() {
			log.Fatalf("Error connecting to ES Server: %s", res.String())
		}
		// Deserialize the response into a map.
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}

		// Print client and server version numbers.
		log.Printf("Connected to Elasticsearch Server")
		log.Printf("ES Client Version: %s", elasticsearch.Version)
		log.Printf("ES Server Version: %s", r["version"].(map[string]interface{})["number"])
	} else {
		log.Println("Analytics Server is already initiated.")
	}
}

func GetEsClient() *elasticsearch.Client {
	if esClient == nil {
		Init()
	}
	return esClient
}
