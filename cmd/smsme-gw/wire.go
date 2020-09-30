//+build wireinject

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

var awsSet = wire.NewSet(
	aws.NewSession,
)

var iftttSet = wire.NewSet(
	ifttt.NewIftttSMSGateway,
)

var sqsSet = wire.NewSet(
	sqs.New,
)

var handlerSet = wire.NewSet(
	console.NewIssueRecorder,
	sms.NewSMSHandler,
)

func InitializeSQSHandler() (inbound.SQSEventHandler, error) {
	wire.Build(awsSet, sqsSet, iftttSet, handlerSet)

	return nil, nil
}
