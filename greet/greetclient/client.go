package main

import (
	"fmt"

	"github.com/maria-robobug/grpc-go-course/greet/greetpb"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// WithInsecure should be temporary, needs certificates for secure connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	logrus.Infof("Created client: %f", c)
}
