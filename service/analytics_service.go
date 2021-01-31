package service

import (
	"context"
	"log"
	"softgo-analytics/analytics"
	analyticsPb "softgo-analytics/proto/analytics"
)

type AnalyticsServer struct {
	analyticsPb.UnimplementedAnalyticsServer
}

// Ping
func (*AnalyticsServer) Ping(ctx context.Context, req *analyticsPb.PingRequest) (*analyticsPb.PingResponse, error) {
	log.Println("Incoming Ping Request.")
	log.Println("Ping -> ", req.Ping)
	response := "pong from analytics server"
	res := &analyticsPb.PingResponse{
		Pong: response,
	}
	return res, nil
}

func (*AnalyticsServer) PeriodicSalesAmount(ctx context.Context, req *analyticsPb.PeriodicSalesAmountRequest) (*analyticsPb.PeriodicSalesAmountResponse, error) {
	log.Println("Incoming Periodic Sales Amount Request.")

	data := analytics.GetPeriodicSalesAmount(req.ChannelName, req.StoreId, req.Interval, req.From, req.To)

	var salesData []*analyticsPb.SalesData;
	for _, v := range data["sales"].(map[string]interface{})["buckets"].([]interface{}) {
		salesData = append(salesData, &analyticsPb.SalesData{
			Timestamp:   v.(map[string]interface{})["key"].(float64),
			BillCount:   v.(map[string]interface{})["doc_count"].(float64),
			TotalAmount: v.(map[string]interface{})["total"].(map[string]interface{})["value"].(float64),
		})
	}

	res := &analyticsPb.PeriodicSalesAmountResponse{
		Success:   true,
		Total:     data["totalSales"].(map[string]interface{})["value"].(float64),
		SalesData: salesData,
	}
	return res, nil
}

func (*AnalyticsServer) TopSoldItems(ctx context.Context, req *analyticsPb.TopSoldItemsRequest) (*analyticsPb.TopSoldItemsResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetTopSoldItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var soldItemsData []*analyticsPb.SoldItemsData;
	for _, v := range data["top_sales"].(map[string]interface{})["buckets"].([]interface{}) {
		soldItemsData = append(soldItemsData, &analyticsPb.SoldItemsData{
			Key:   v.(map[string]interface{})["key"].(string),
			Count: v.(map[string]interface{})["doc_count"].(float64),
		})
	}
	
	res := &analyticsPb.TopSoldItemsResponse{
		Success:       true,
		SoldItemsData: soldItemsData,
	}
	return res, nil
}

func (*AnalyticsServer) LeastSoldItems(ctx context.Context, req *analyticsPb.LeastSoldItemsRequest) (*analyticsPb.LeastSoldItemsResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetLeastSoldItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var soldItemsData []*analyticsPb.SoldItemsData;
	for _, v := range data["top_sales"].(map[string]interface{})["buckets"].([]interface{}) {
		soldItemsData = append(soldItemsData, &analyticsPb.SoldItemsData{
			Key:   v.(map[string]interface{})["key"].(string),
			Count: v.(map[string]interface{})["doc_count"].(float64),
		})
	}

	res := &analyticsPb.LeastSoldItemsResponse{
		Success:       true,
		SoldItemsData: soldItemsData,
	}
	return res, nil
}

func (*AnalyticsServer) TotalSalesAmount(ctx context.Context, req *analyticsPb.TotalSalesAmountRequest) (*analyticsPb.TotalSalesAmountResponse, error) {
	log.Println("Incoming Total Sales Amount Request.")

	amount := analytics.GetTotalSalesAmount(req.ChannelName, req.StoreId, req.From, req.To)
	res := &analyticsPb.TotalSalesAmountResponse{
		Success: true,
		Amount:  amount,
	}
	return res, nil
}
