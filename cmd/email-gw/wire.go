//+build wireinject

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

var awsSet = wire.NewSet(
	aws.NewSession,
)

var sesSet = wire.NewSet(
	ses.NewSESEmailGateway,
)

var sqsSet = wire.NewSet(
	sqs.New,
)

var handlerSet = wire.NewSet(
	console.NewIssueRecorder,
	email.NewSQSHandler,
)

func InitializeSQSHandler() (inbound.SQSEventHandler, error) {
	wire.Build(awsSet, sqsSet, sesSet, handlerSet)

	return nil, nil
}
