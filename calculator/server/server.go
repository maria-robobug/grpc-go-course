package main

import (
	"context"
	"net"

	"github.com/maria-robobug/grpc-go-course/calculator/calcpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type service struct{}

func (*service) Calculate(ctx context.Context, req *calcpb.CalculatorRequest) (*calcpb.CalculatorResponse, error) {
	log.Infof("Calculate function invoked with: %v\n", req)

	firstNum := req.GetNumbers().GetFirstNum()
	secondNum := req.GetNumbers().GetSecondNum()

	resp := &calcpb.CalculatorResponse{
		Result: firstNum + secondNum,
	}

	return resp, nil
}

func main() {
	log.Infoln("Started RPC calculator server")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &service{})

	if err := s.Serve(list); err != nil {
		log.Errorf("failed to serve: %v", err)
	}
}
