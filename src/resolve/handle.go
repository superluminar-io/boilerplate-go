package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

	table := os.Getenv("DYNAMODB_TABLE_NAME")
	id := request.PathParameters["id"]

	s := session.Must(session.NewSession())
	client := dynamodb.New(s)
	result, err := client.GetItem(
		&dynamodb.GetItemInput{
			TableName: aws.String(table),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {S: aws.String(id)},
			},
		},
	)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to access data",
		}, nil
	}

	if result.Item == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Unknown ID",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 302,
		Headers: map[string]string{
			"Location": *result.Item["url"].S,
		},
	}, nil

}
