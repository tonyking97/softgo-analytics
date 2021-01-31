package service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	analyticsPb "softgo-analytics/proto/analytics"
)

func GrpcInit() {
	go startGRPCServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	analyticsPb.RegisterAnalyticsServer(s, &AnalyticsServer{})

	reflection.Register(s)
	log.Println("Starting gRPC Service...")
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to start gRPC Service.", err)
	}
}
