// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/finfech/notiplat-gateway/inbound"
	"github.com/finfech/notiplat-gateway/inbound/email"
	"github.com/finfech/notiplat-gateway/infra/aws"
	"github.com/finfech/notiplat-gateway/infra/console"
	"github.com/finfech/notiplat-gateway/infra/ses"
	"github.com/finfech/notiplat-gateway/infra/sqs"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeSQSHandler() (inbound.SQSEventHandler, error) {
	session, err := aws.NewSession()
	if err != nil {
		return nil, err
	}
	gateway, err := ses.NewSESEmailGateway(session)
	if err != nil {
		return nil, err
	}
	recorder := console.NewIssueRecorder()
	sqsSQS := sqs.New(session)
	sqsEventHandler, err := email.NewSQSHandler(gateway, recorder, sqsSQS)
	if err != nil {
		return nil, err
	}
	return sqsEventHandler, nil
}

// wire.go:

var awsSet = wire.NewSet(aws.NewSession)

var sesSet = wire.NewSet(ses.NewSESEmailGateway)

var sqsSet = wire.NewSet(sqs.New)

var handlerSet = wire.NewSet(console.NewIssueRecorder, email.NewSQSHandler)
