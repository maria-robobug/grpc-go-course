package main

import (
	"context"
	"fmt"

	"github.com/maria-robobug/grpc-go-course/greet/greetpb"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	// WithInsecure should be temporary, needs certificates for secure connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Errorf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Infoln("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Maria",
			LastName:  "Shodunke",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Errorf("error while calling Greet RPC: %v", err)
	}

	log.Infof("response from Greet: %v", resp.Result)
}
