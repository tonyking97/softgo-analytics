package main

import (
	"softgo-analytics/es"
	"softgo-analytics/service"
)

func main() {
	done := make(chan bool, 1)

	es.Init()
	//analytics.GetBills()
	//for{
	//	start := time.Now()
	//	amount := analytics.GetSalesAmount("softgo")
	//	elapsed := time.Since(start)
	//	log.Println("Time Taken - ", elapsed)
	//	time.Sleep(1 * time.Second)
	//	log.Printf("%f\n",amount)
	//}

	service.GrpcInit()

	<-done
}
