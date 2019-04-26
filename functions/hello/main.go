package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Id    string `json:"id"`
	Value string
}

func greet(in *Request) (events.APIGatewayProxyResponse, error) {
	bytes, _ := json.Marshal(Response{
		Id:    time.Now().String(),
		Value: fmt.Sprintf("Hello %s", in.Name),
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(bytes),
	}, nil
}

func main() {
	lambda.Start(greet)
}
