package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// New godoc
func New(session *session.Session) *sqs.SQS {
	return sqs.New(session)
}
