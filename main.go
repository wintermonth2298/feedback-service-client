package main

import (
	"context"
	"log"

	fs "github.com/wintermonth2298/feedback-service/api/feedback-service"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial(): %v", err)
	}
	defer conn.Close()

	service := fs.NewFeedbackServiceClient(conn)

	response, err := service.GetFeedback(context.Background(), &fs.GetFeedbackRequest{Id: 1})

	if err != nil {
		log.Fatalf("GetFeedbackRequest(): %v", err)
	}
	log.Println(response)
}
