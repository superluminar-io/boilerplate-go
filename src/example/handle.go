package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"os"
)

type event struct {
	ShouldFail bool `json:"should_fail"`
}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf(
		"Function of %s-%s invoked",
		os.Getenv("PREFIX"),
		os.Getenv("PROJECT"),
	)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Example String",
	}, nil
}
