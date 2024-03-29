// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/finfech/notiplat-gateway/inbound"
	"github.com/finfech/notiplat-gateway/inbound/sms"
	"github.com/finfech/notiplat-gateway/infra/aws"
	"github.com/finfech/notiplat-gateway/infra/console"
	"github.com/finfech/notiplat-gateway/infra/ifttt"
	"github.com/finfech/notiplat-gateway/infra/sqs"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeSQSHandler() (inbound.SQSEventHandler, error) {
	gateway, err := ifttt.NewIftttSMSGateway()
	if err != nil {
		return nil, err
	}
	recorder := console.NewIssueRecorder()
	session, err := aws.NewSession()
	if err != nil {
		return nil, err
	}
	sqsSQS := sqs.New(session)
	sqsEventHandler, err := sms.NewSMSHandler(gateway, recorder, sqsSQS)
	if err != nil {
		return nil, err
	}
	return sqsEventHandler, nil
}

// wire.go:

var awsSet = wire.NewSet(aws.NewSession)

var iftttSet = wire.NewSet(ifttt.NewIftttSMSGateway)

var sqsSet = wire.NewSet(sqs.New)

var handlerSet = wire.NewSet(console.NewIssueRecorder, sms.NewSMSHandler)
