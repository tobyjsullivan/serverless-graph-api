package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"log"
)

type ApiEvent struct {
	Body       string `json:"body"`
	HttpMethod string `json:"httpMethod"`
	Resource   string `json:"resource"`
}

type ApiResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            interface{}       `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func generateApiResponse(statusCode int, headers map[string]string, body interface{}) (ApiResponse, error) {
	if content, err := json.Marshal(body); err != nil {
		return ApiResponse{}, err
	} else {
		return ApiResponse{
			StatusCode:      statusCode,
			Body:            string(content),
			Headers:         headers,
			IsBase64Encoded: false,
		}, nil
	}
}

func HandleLambdaEvent(event ApiEvent) (ApiResponse, error) {
	name := event.Body
	resp, err := generateApiResponse(200, map[string]string{
		"x-method":   event.HttpMethod,
		"x-resource": event.Resource,
	}, MyResponse{
		Message: fmt.Sprintf("Received: %s!", name),
	})
	if err != nil {
		log.Println("Encountered error in execution:", err.Error())
		return ApiResponse{}, err
	}

	serResp, err := json.Marshal(resp)
	if err != nil {
		panic("Unexpected error during response serialization: " + err.Error())
	}

	log.Println("Returning this response: " + string(serResp))
	return resp, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
