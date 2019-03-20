package main

import (
	"context"

	"github.com/maria-robobug/grpc-go-course/calculator/calcpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Infoln("Started RPC calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Errorf("could not connect: %v", err)
	}

	defer cc.Close()

	c := calcpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calcpb.CalculatorServiceClient) {
	log.Infoln("Starting to do a Unary RPC")

	req := &calcpb.CalculatorRequest{
		Numbers: &calcpb.Numbers{
			FirstNum:  4,
			SecondNum: 26,
		},
	}

	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Errorf("failed to calculate numbers: %v", err)
	}

	log.Infof("Calculator: %v + %v = %v", req.Numbers.FirstNum, req.Numbers.SecondNum, resp.Result)
}
