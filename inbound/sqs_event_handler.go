package inbound

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// SQSEventHandler godoc
type SQSEventHandler func(ctx context.Context, sqsEvent events.SQSEvent) error
