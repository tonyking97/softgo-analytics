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

func (*AnalyticsServer) AddStore(ctx context.Context, req *analyticsPb.AddStoreRequest) (*analyticsPb.AddStoreResponse, error) {
	log.Println("Incoming Add Store Request.")

	err := analytics.AddStore(req.ChannelName)
	if err != nil {
		return &analyticsPb.AddStoreResponse{
			Success: false,
			Message: "failed to create store",
		}, nil
	} else {
		return &analyticsPb.AddStoreResponse{
			Success: true,
		}, nil
	}
}

func (*AnalyticsServer) DeleteStore(ctx context.Context, req *analyticsPb.DeleteStoreRequest) (*analyticsPb.DeleteStoreResponse, error) {
	log.Println("Incoming Delete Store Request.")

	err := analytics.DeleteStore(req.ChannelName)
	if err != nil {
		return &analyticsPb.DeleteStoreResponse{
			Success: false,
			Message: "failed to delete store",
		}, nil
	} else {
		return &analyticsPb.DeleteStoreResponse{
			Success: true,
		}, nil
	}
}

func (*AnalyticsServer) PeriodicSalesAmount(ctx context.Context, req *analyticsPb.PeriodicSalesAmountRequest) (*analyticsPb.PeriodicSalesAmountResponse, error) {
	log.Println("Incoming Periodic Sales Amount Request.")

	data := analytics.GetPeriodicSalesAmount(req.ChannelName, req.StoreId, req.Interval, req.From, req.To)

	var salesData []*analyticsPb.SalesData
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

func (*AnalyticsServer) PeriodicStoreSalesAmount(ctx context.Context, req *analyticsPb.PeriodicStoreSalesAmountRequest) (*analyticsPb.PeriodicStoreSalesAmountResponse, error) {
	log.Println("Incoming Periodic Store Sales Amount Request.")

	data := analytics.GetPeriodicStoreSalesAmount(req.ChannelName, req.Interval, req.StoreCount, req.From, req.To)

	var storeSalesData []*analyticsPb.StoreSalesData

	for _, storeData := range data["sales_data"].(map[string]interface{})["buckets"].([]interface{}) {
		var salesData []*analyticsPb.SalesData
		for _, v := range storeData.(map[string]interface{})["sales"].(map[string]interface{})["buckets"].([]interface{}) {
			salesData = append(salesData, &analyticsPb.SalesData{
				Timestamp:   v.(map[string]interface{})["key"].(float64),
				BillCount:   v.(map[string]interface{})["doc_count"].(float64),
				TotalAmount: v.(map[string]interface{})["total"].(map[string]interface{})["value"].(float64),
			})
		}
		storeSalesData = append(storeSalesData, &analyticsPb.StoreSalesData{
			StoreId:   storeData.(map[string]interface{})["key"].(string),
			BillCount: storeData.(map[string]interface{})["doc_count"].(float64),
			Total:     storeData.(map[string]interface{})["totalSales"].(map[string]interface{})["value"].(float64),
			SalesData: salesData,
		})
	}

	res := &analyticsPb.PeriodicStoreSalesAmountResponse{
		Success:        true,
		StoreSalesData: storeSalesData,
	}
	return res, nil
}

func (*AnalyticsServer) TopSoldItems(ctx context.Context, req *analyticsPb.TopSoldItemsRequest) (*analyticsPb.TopSoldItemsResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetTopSoldItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var soldItemsData []*analyticsPb.SoldItemsData
	for _, v := range data["sales_data"].(map[string]interface{})["top_sold_items"].(map[string]interface{})["buckets"].([]interface{}) {
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
	log.Println("Incoming Least Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetLeastSoldItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var soldItemsData []*analyticsPb.SoldItemsData
	for _, v := range data["sales_data"].(map[string]interface{})["least_sold_items"].(map[string]interface{})["buckets"].([]interface{}) {
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

func (*AnalyticsServer) TopRevenueItems(ctx context.Context, req *analyticsPb.TopRevenueItemsRequest) (*analyticsPb.TopRevenueItemsResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetTopRevenueItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var revenueItemsData []*analyticsPb.RevenueItemsData
	for _, v := range data["sales_data"].(map[string]interface{})["top_revenue_items"].(map[string]interface{})["buckets"].([]interface{}) {
		revenueItemsData = append(revenueItemsData, &analyticsPb.RevenueItemsData{
			Key:             v.(map[string]interface{})["key"].(string),
			BillCount:       v.(map[string]interface{})["doc_count"].(float64),
			TotalSalesPrice: v.(map[string]interface{})["total_sales_price"].(map[string]interface{})["value"].(float64),
		})
	}

	res := &analyticsPb.TopRevenueItemsResponse{
		Success:          true,
		RevenueItemsData: revenueItemsData,
	}
	return res, nil
}

func (*AnalyticsServer) LeastRevenueItems(ctx context.Context, req *analyticsPb.LeastRevenueItemsRequest) (*analyticsPb.LeastRevenueItemsResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetLeastRevenueItems(req.ChannelName, req.StoreId, req.From, req.To, req.Count)

	var revenueItemsData []*analyticsPb.RevenueItemsData
	for _, v := range data["sales_data"].(map[string]interface{})["top_revenue_items"].(map[string]interface{})["buckets"].([]interface{}) {
		revenueItemsData = append(revenueItemsData, &analyticsPb.RevenueItemsData{
			Key:             v.(map[string]interface{})["key"].(string),
			BillCount:       v.(map[string]interface{})["doc_count"].(float64),
			TotalSalesPrice: v.(map[string]interface{})["total_sales_price"].(map[string]interface{})["value"].(float64),
		})
	}

	res := &analyticsPb.LeastRevenueItemsResponse{
		Success:          true,
		RevenueItemsData: revenueItemsData,
	}
	return res, nil
}

func (*AnalyticsServer) StoreRevenueData(ctx context.Context, req *analyticsPb.StoreRevenueDataRequest) (*analyticsPb.StoreRevenueDataResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetStoreRevenueData(req.ChannelName, req.StoreId, req.From, req.To, req.Count, req.Ascending)

	var storeRevenueData []*analyticsPb.StoreRevenueData
	for _, v := range data["sales_data"].(map[string]interface{})["buckets"].([]interface{}) {
		storeRevenueData = append(storeRevenueData, &analyticsPb.StoreRevenueData{
			Key:          v.(map[string]interface{})["key"].(string),
			BillCount:    v.(map[string]interface{})["doc_count"].(float64),
			TotalRevenue: v.(map[string]interface{})["total_sales"].(map[string]interface{})["value"].(float64),
		})
	}

	res := &analyticsPb.StoreRevenueDataResponse{
		Success:          true,
		StoreRevenueData: storeRevenueData,
	}
	return res, nil
}

func (*AnalyticsServer) StoreSalesData(ctx context.Context, req *analyticsPb.StoreSalesDataRequest) (*analyticsPb.StoreSalesDataResponse, error) {
	log.Println("Incoming Top Sold Items Request.")

	if req.Count <= 0{
		req.Count = 1
	}
	data := analytics.GetStoreSalesData(req.ChannelName, req.StoreId, req.From, req.To, req.Count, req.Ascending)

	var storeBillData []*analyticsPb.StoreBillData
	for _, v := range data["sales_data"].(map[string]interface{})["buckets"].([]interface{}) {
		storeBillData = append(storeBillData, &analyticsPb.StoreBillData{
			Key:          v.(map[string]interface{})["key"].(string),
			BillCount:    v.(map[string]interface{})["doc_count"].(float64),
		})
	}

	res := &analyticsPb.StoreSalesDataResponse{
		Success:       true,
		StoreBillData: storeBillData,
	}
	return res, nil
}

func (*AnalyticsServer) AverageSalesAmount(ctx context.Context, req *analyticsPb.AverageSalesAmountRequest) (*analyticsPb.AverageSalesAmountResponse, error) {
	log.Println("Incoming Average Sales Amount Request.")

	amount := analytics.GetAverageSalesAmount(req.ChannelName, req.StoreId, req.From, req.To)
	res := &analyticsPb.AverageSalesAmountResponse{
		Success: true,
		Amount:  amount,
	}
	return res, nil
}

func (*AnalyticsServer) AverageSoldQuantity(ctx context.Context, req *analyticsPb.AverageSoldQuantityRequest) (*analyticsPb.AverageSoldQuantityResponse, error) {
	log.Println("Incoming Average Sales Amount Request.")

	quantity := analytics.GetAverageSoldQuantity(req.ChannelName, req.StoreId, req.From, req.To)
	res := &analyticsPb.AverageSoldQuantityResponse{
		Success:  true,
		Quantity: quantity,
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
