package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func HandleRequest(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is old!", event.Name)}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
