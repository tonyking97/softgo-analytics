package analytics

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"softgo-analytics/es"
	"strings"
	"time"
)

// TODO Handle error for invalid input params

func AddStore(channelName string) error {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var buf bytes.Buffer
	mappings := map[string]interface{}{
		"mappings" : map[string]interface{}{
			"properties" : map[string]interface{}{
				"timestamp" : map[string]interface{}{
					"type" : "date",
					"format": "epoch_second",
				},
				"dishes" : map[string]interface{}{
					"type" : "nested",
					"properties" : map[string]interface{}{
						"name" : map[string]interface{}{
							"type": "text",
							"fielddata": true,
						},
						"price" : map[string]interface{}{
							"type": "long",
						},
					},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(mappings); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := esClient.Indices.Create(
		indexName,
		esClient.Indices.Create.WithContext(context.Background()),
		esClient.Indices.Create.WithBody(&buf),
	)

	if err != nil {
		log.Println("Error getting response: %s", err)
		return errors.New("Internal Server Error")
	}

	defer res.Body.Close()

	if res.IsError() {
		var resError map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resError); err != nil {
			log.Println("Error parsing the response body: %s", err)
			return errors.New("Error parsing the response")
		} else {
			// Print the response status and resError information.
			log.Println("[",res.Status(),"]",
				"type:",
				resError["error"].(map[string]interface{})["type"],
				"reason:",
				resError["error"].(map[string]interface{})["reason"],
			)
			if resError["error"].(map[string]interface{})["type"] == "resource_already_exists_exception" {
				return errors.New("Index Already Exists.")
			} else {
				return errors.New("Failed to create index.")
			}
		}
	}

	return nil
}

func DeleteStore(channelName string) error {
	esClient := es.GetEsClient()
	indexName := []string{channelName + "_bill"}

	var r map[string]interface{}

	res, err := esClient.Indices.Delete(
		indexName,
		esClient.Indices.Delete.WithContext(context.Background()),
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

	log.Println(r)
	return nil
}

func AddBill(channelName string, billData []byte) error {
	t := time.Now()
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"
	_, err := esClient.Index(indexName, bytes.NewReader(billData))
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return errors.New("Internal Server Error")
	}
	log.Println("Added Bill Data to Index. Time Taken ->", time.Since(t))
	return nil
}

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

func GetPeriodicStoreSalesAmount(channelName, interval string, storeCount, from, to int64) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	//Default interval as month
	if interval=="" {
		interval = "month"
	}

	//Default store count
	if storeCount==0 {
		storeCount = 50
	}

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs" : map[string]interface{}{
			"sales_data" : map[string]interface{}{
				"terms" : map[string]interface{}{
					"field" : "store.keyword",
					"order" : map[string]interface{}{
						"_key" : "asc",
					},
					"size" : storeCount,
				},
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
			},
		},
	}

	var filterQuery []map[string]interface{}

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

func GetPeriodicSalesAmount(channelName, storeId, interval string, from, to int64) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	//Default interval as month
	if interval=="" {
		interval = "month"
	}

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
			"sales_data": map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "dishes",
				},
				"aggs": map[string]interface{}{
					"top_sold_items": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "dishes.name",
							"size": count,
						},
					},
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
			"sales_data": map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "dishes",
				},
				"aggs": map[string]interface{}{
					"least_sold_items": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "dishes.name",
							"order": map[string]interface{}{
								"_count": "asc",
							},
							"size": count,
						},
					},
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

func GetTopRevenueItems(channelName, storeId string, from, to int64, count int32) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales_data": map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "dishes",
				},
				"aggs": map[string]interface{}{
					"top_revenue_items": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "dishes.name",
							"order": map[string]interface{}{
								"total_sales_price": "desc",
							},
							"size": count,
						},
						"aggs" : map[string]interface{}{
							"total_sales_price" : map[string]interface{}{
								"sum" : map[string]interface{}{
									"field" : "dishes.price",
								},
							},
						},
					},
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

func GetLeastRevenueItems(channelName, storeId string, from, to int64, count int32) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales_data": map[string]interface{}{
				"nested": map[string]interface{}{
					"path": "dishes",
				},
				"aggs": map[string]interface{}{
					"top_revenue_items": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "dishes.name",
							"order": map[string]interface{}{
								"total_sales_price": "asc",
							},
							"size": count,
						},
						"aggs" : map[string]interface{}{
							"total_sales_price" : map[string]interface{}{
								"sum" : map[string]interface{}{
									"field" : "dishes.price",
								},
							},
						},
					},
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

func GetStoreRevenueData(channelName, storeId string, from, to int64, count int32, ascending bool) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	//Default sorting in descending order
	sort := "desc"
	if ascending == true {
		sort = "asc"
	}

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales_data": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "store.keyword",
					"order": map[string]interface{}{
						"total_sales": sort,
					},
					"size": count,
				},
				"aggs": map[string]interface{}{
					"total_sales" : map[string]interface{}{
						"sum" : map[string]interface{}{
							"field" : "total",
						},
					},
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

func GetStoreSalesData(channelName, storeId string, from, to int64, count int32, ascending bool) map[string]interface{} {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	//Default sorting in descending order
	sort := "desc"
	if ascending == true {
		sort = "asc"
	}

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales_data": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "store.keyword",
					"order": map[string]interface{}{
						"_count": sort,
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

func GetAverageSalesAmount(channelName, storeId string, from, to int64) float64 {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"sales": map[string]interface{}{
				"avg": map[string]interface{}{
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

	salesAmount:= r["aggregations"].(map[string]interface{})["sales"].(map[string]interface{})["value"].(float64)
	return salesAmount
}

func GetAverageSoldQuantity(channelName, storeId string, from, to int64) float64 {
	esClient := es.GetEsClient()
	indexName := channelName + "_bill"

	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"aggs": map[string]interface{}{
			"average_sold_quantity": map[string]interface{}{
				"avg": map[string]interface{}{
					"field": "total_quantity",
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

	averageQuantity := r["aggregations"].(map[string]interface{})["average_sold_quantity"].(map[string]interface{})["value"].(float64)
	return averageQuantity
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
