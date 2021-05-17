package main

import (
	"softgo-analytics/es"
	"softgo-analytics/queue"
	"softgo-analytics/service"
)

func main() {
	done := make(chan bool, 1)

	es.Init()
	queue.StartQueueProcessor()
	service.GrpcInit()

	<-done
}
