package analytics

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"softgo-analytics/es"
	"strings"
)


func GetBills() {
	esClient := es.GetEsClient()

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"store": "store_1",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex("softgo_bill"),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

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

	log.Println(strings.Repeat("~", 37))
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
	//log.Println(r)
	//Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}

func GetPeriodicSalesAmount(channelName, storeId, interval string, from, to int64) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales": map[string]interface{}{
				"date_histogram": map[string]interface{}{
					"field": "timestamp",
					"calendar_interval": interval,
				},
				"aggs": map[string]interface{}{
					"total": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "total",
						},
					},
				},
			},
			"totalSales": map[string]interface{}{
				"sum_bucket": map[string]interface{}{
					"buckets_path": "sales>total",
				},
			},
		},
	}

	var filterQuery []map[string]interface{}

	if storeId != "" {
		filterQuery = append(filterQuery, map[string]interface{}{
			"term": map[string]interface{}{
				"store": storeId,
			},
		})
	}

	if from > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": from,
				},
			},
		})
	}

	if to > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"lt": to,
				},
			},
		})
	}

	query["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"filter": filterQuery,
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(false),
		esClient.Search.WithSize(0),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var resError map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resError); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and resError information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				resError["resError"].(map[string]interface{})["type"],
				resError["resError"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return r["aggregations"].(map[string]interface{})
}

func GetTopSoldItems(channelName, storeId string, from, to int64, count int32) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"top_sales": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "dishes.name",
					"size": count,
				},
			},
		},
	}

	var filterQuery []map[string]interface{}

	if storeId != "" {
		filterQuery = append(filterQuery, map[string]interface{}{
			"term": map[string]interface{}{
				"store": storeId,
			},
		})
	}

	if from > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": from,
				},
			},
		})
	}

	if to > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"lt": to,
				},
			},
		})
	}

	query["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"filter": filterQuery,
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(false),
		esClient.Search.WithSize(0),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var resError map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resError); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and resError information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				resError["resError"].(map[string]interface{})["type"],
				resError["resError"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return r["aggregations"].(map[string]interface{})
}

func GetLeastSoldItems(channelName, storeId string, from, to int64, count int32) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"top_sales": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "dishes.name",
					"order": map[string]interface{}{
						"_count": "asc",
					},
					"size": count,
				},
			},
		},
	}

	var filterQuery []map[string]interface{}

	if storeId != "" {
		filterQuery = append(filterQuery, map[string]interface{}{
			"term": map[string]interface{}{
				"store": storeId,
			},
		})
	}

	if from > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": from,
				},
			},
		})
	}

	if to > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"lt": to,
				},
			},
		})
	}

	query["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"filter": filterQuery,
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(false),
		esClient.Search.WithSize(0),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var resError map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resError); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and resError information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				resError["resError"].(map[string]interface{})["type"],
				resError["resError"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	return r["aggregations"].(map[string]interface{})
}

func GetTotalSalesAmount(channelName, storeId string, from, to int64) float64 {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales_stats": map[string]interface{}{
				"sum": map[string]interface{}{
					"field": "total",
				},
			},
		},
	}

	var filterQuery []map[string]interface{}

	if storeId != "" {
		filterQuery = append(filterQuery, map[string]interface{}{
			"term": map[string]interface{}{
				"store": storeId,
			},
		})
	}

	if from > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": from,
				},
			},
		})
	}

	if to > 0 {
		filterQuery = append(filterQuery, map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"lt": to,
				},
			},
		})
	}

	query["query"] = map[string]interface{}{
		"bool": map[string]interface{}{
			"filter": filterQuery,
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(indexName),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithSize(0),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var resError map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resError); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and resError information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				resError["resError"].(map[string]interface{})["type"],
				resError["resError"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	salesAmount:= r["aggregations"].(map[string]interface{})["sales_stats"].(map[string]interface{})["value"].(float64)
	return salesAmount
}
