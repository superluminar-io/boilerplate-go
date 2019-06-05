package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/url"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type event struct {
	ShouldFail bool `json:"should_fail"`
}

func handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	s := session.Must(session.NewSession())
	client := dynamodb.New(s)

	var data map[string]string
	err := json.Unmarshal([]byte(request.Body), &data)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	url, _ := data["url"]

	// Error Handling

	id, err := shorten(url)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	_, err = client.PutItem(
		&dynamodb.PutItemInput{
			TableName: aws.String(table),
			Item: map[string]*dynamodb.AttributeValue{
				"id":  {S: aws.String(id)},
				"url": {S: aws.String(url)},
			},
		},
	)
	body := fmt.Sprintf("Created short url: %s/%s/%s", request.Headers["Host"], "Prod", id)
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       body,
	}, nil
}

func shorten(u string) (string, error) {
	if _, err := url.ParseRequestURI(u); err != nil {
		return "", err
	}

	hash := fnv.New64a()
	if _, err := hash.Write([]byte(u)); err != nil {
		return "", err
	}

	return strconv.FormatUint(hash.Sum64(), 36), nil
}
