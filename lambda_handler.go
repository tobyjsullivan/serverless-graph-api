package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"log"
	"github.com/graphql-go/graphql"
	"net/http"
	"errors"
)

var (
	schema graphql.Schema
)

type ApiEvent struct {
	Body       string `json:"body"`
	HttpMethod string `json:"httpMethod"`
	Path       string `json:"path"`
}

type QueryRequest struct {
	Query string `json:"query"`
}

type ApiResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            interface{}       `json:"body"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

type ResponseBody struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}

func generateApiResponse(statusCode int, body interface{}) (ApiResponse, error) {
	if content, err := json.Marshal(body); err != nil {
		return ApiResponse{}, err
	} else {
		return ApiResponse{
			StatusCode:      statusCode,
			Body:            string(content),
			Headers:         map[string]string{},
			IsBase64Encoded: false,
		}, nil
	}
}

func HandleLambdaEvent(event ApiEvent) (ApiResponse, error) {
	log.Printf("Request path: %s\n", event.Path)
	log.Printf("Request method: %s\n", event.HttpMethod)
	log.Printf("Request body: %s\n", event.Body)

	if event.Path == "/query" && event.HttpMethod == http.MethodPost {
		var queryRequest QueryRequest
		err := json.Unmarshal([]byte(event.Body), &queryRequest)
		if err != nil {
			return generateApiResponse(http.StatusBadRequest, ResponseBody{Errors: []string{err.Error()}})
		}

		return handleGraphQuery(&queryRequest)
	}

	return ApiResponse{}, errors.New(fmt.Sprintf("Unexpected query: %s %s", event.HttpMethod, event.Path))
}

func handleGraphQuery(query *QueryRequest) (ApiResponse, error) {
	params := graphql.Params{Schema: schema, RequestString: query.Query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return generateApiResponse(http.StatusBadRequest, r)
	}

	return generateApiResponse(http.StatusOK, r)
}

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema: %v", err)
	}

	lambda.Start(HandleLambdaEvent)
}
