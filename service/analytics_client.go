package service

import (
	"context"
	"log"
	"softgo-analytics/models"
	analyticsClientPb "softgo-analytics/proto/analytics_client"
	"softgo-analytics/queue"
	"time"
)

type AnalyticsClient struct {
	analyticsClientPb.UnimplementedAnalyticsClientServer
}

func (*AnalyticsClient) HealthCheck(ctx context.Context, req *analyticsClientPb.HealthCheckRequest) (*analyticsClientPb.HealthCheckResponse, error) {
	log.Println("Incoming HealthCheck Request.")
	res := &analyticsClientPb.HealthCheckResponse{}
	return res, nil
}

func (*AnalyticsClient) AddBill(ctx context.Context, req *analyticsClientPb.AddBillRequest) (*analyticsClientPb.AddBillResponse, error) {
	log.Println("Add Bill Request.")
	t := time.Now()
	var dishes []models.DishBill
	for _, v := range req.Dishes{
		dishes = append(dishes, models.DishBill{
			Name:  v.Name,
			Price: v.Price,
		})
	}
	err := queue.AddBillTask(req.ChannelName, req.StoreId, req.CustomerName, req.Timestamp, dishes, req.TotalQuantity, req.Total)
	if err != nil {
		log.Println("Error enqueue bill data ->", err)
		return &analyticsClientPb.AddBillResponse{Success: false}, nil
	}
	log.Println("Time Taken -->", time.Since(t))
	res := &analyticsClientPb.AddBillResponse{Success: true}
	return res, nil
}
