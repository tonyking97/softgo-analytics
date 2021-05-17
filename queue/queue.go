package queue

import (
	"github.com/hibiken/asynq"
	"log"
)

const redisAddr = "127.0.0.1:6379"
var queueClient *asynq.Client = nil

func InitClient() {
	if queueClient == nil {
		redisClient := asynq.RedisClientOpt{Addr: redisAddr}
		queueClient = asynq.NewClient(redisClient)
	}
}

func GetQueueClient() *asynq.Client {
	if queueClient == nil {
		InitClient()
	}
	return queueClient
}

func StartQueueProcessor(){
	go startQueueServer()
}

func startQueueServer(){
	redisClient := asynq.RedisClientOpt{Addr: redisAddr}
	queueServer := asynq.NewServer(redisClient, asynq.Config{
		Concurrency:         10,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(BillTaskName, ProcessBillTask)

	if err := queueServer.Run(mux); err != nil {
		log.Fatal(err)
	}
}
