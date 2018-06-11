package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(ctx context.Context, event MyEvent) (MyResponse, error) {
	name := event.Name
	return MyResponse{
		Message: fmt.Sprintf("Hello, friend %s!", name),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
