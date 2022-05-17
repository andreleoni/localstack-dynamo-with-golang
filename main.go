package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	ClusterID string `json:"cluster_id"`
	ProductID string `json:"product_id"`
	Priority  string `json:"priority"`
}

func main() {
	endpoint := "http://localhost:4566"
	region := "us-east-1"

	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
	})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// monta o item que será salvo
	item := Item{
		ClusterID: "baldao",
		ProductID: "2",
		Priority:  "ALTERADO DE NOVO",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	tableName := "recommendations"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Criado")
	}

	// Avoid record update
	avoidRecordUpdateInput := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(cluster_id)"),
	}

	_, err = svc.PutItem(avoidRecordUpdateInput)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Criado")
	}
}
