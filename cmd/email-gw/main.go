package main

import "github.com/aws/aws-lambda-go/lambda"

func main() {
	handler, err := InitializeSQSHandler()
	if err != nil {
		panic(err)
	}

	lambda.Start(handler)
}
