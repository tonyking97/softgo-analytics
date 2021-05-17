package queue

import (
	"context"
	"github.com/hibiken/asynq"
	"softgo-analytics/analytics"
	"softgo-analytics/models"
)

const BillTaskName = "analytics:bill"

func AddBillTask(channelName, storeId, customerName string, timestamp int64, dishes []models.DishBill, totalQuantity int32, total int64) error {
	queueClient := GetQueueClient()
	task := asynq.NewTask(BillTaskName, map[string]interface{}{"channel_name": channelName, "customer_name": customerName, "store": storeId, "timestamp": timestamp, "dishes": dishes, "total_quantity": totalQuantity, "total": total})
	_, err := queueClient.Enqueue(task)
	if err != nil {
		return err
	}
	return nil
}

func ProcessBillTask(ctx context.Context, task *asynq.Task) error {
	channelName, _ := task.Payload.GetString("channel_name")
	billData, _ := task.Payload.MarshalJSON()
	err := analytics.AddBill(channelName, billData)
	if err != nil {
		return err
	}
	return nil
}
