//+build wireinject

package main

import (
	"github.com/finfech/notiplat-gateway/inbound"
	"github.com/finfech/notiplat-gateway/inbound/sms"
	"github.com/finfech/notiplat-gateway/infra/aws"
	"github.com/finfech/notiplat-gateway/infra/console"
	"github.com/finfech/notiplat-gateway/infra/sns"
	"github.com/finfech/notiplat-gateway/infra/sqs"
	"github.com/google/wire"
)

var awsSet = wire.NewSet(
	aws.NewSession,
)

var snsSet = wire.NewSet(
	sns.NewSNSSmsGateway,
)

var sqsSet = wire.NewSet(
	sqs.New,
)

var handlerSet = wire.NewSet(
	console.NewIssueRecorder,
	sms.NewSMSHandler,
)

func InitializeSQSHandler() (inbound.SQSEventHandler, error) {
	wire.Build(awsSet, sqsSet, snsSet, handlerSet)

	return nil, nil
}
