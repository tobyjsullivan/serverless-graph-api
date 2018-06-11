package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type ApiEvent struct {
	Body       string `json:"body"`
	HttpMethod string `json:"httpMethod"`
	Resource   string `json:"resource"`
}

type ApiResponse struct {
	StatusCode int `json:"statusCode"`
	Headers []interface{} `json:"headers"`
	Body string `json:"body"`
	IsBase64Encoded bool `json:"isBase64Encoded"`
}

func HandleLambdaEvent(event ApiEvent) (ApiResponse, error) {
	name := event.Body
	return ApiResponse {
		Body: fmt.Sprintf("Received: %s!", name),
		StatusCode: 200,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
