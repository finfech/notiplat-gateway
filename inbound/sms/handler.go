package sms

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/finfech/notiplat-gateway/domain/useCase/issue"
	"github.com/finfech/notiplat-gateway/domain/useCase/sms"
	"github.com/finfech/notiplat-gateway/inbound"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// NewSMSHandler godoc
func NewSMSHandler(gw sms.Gateway,
	issuer issue.Recorder, s *sqs.SQS) (inbound.SQSEventHandler, error) {
	return func(ctx context.Context, sqsEvent events.SQSEvent) error {
		var result *multierror.Error

		processedMessages := []*sqs.DeleteMessageBatchRequestEntry{}

		for _, message := range sqsEvent.Records {
			var req jsonPayload
			if err := json.Unmarshal([]byte(message.Body), &req); err != nil {
				issuer.Record(errors.Wrap(err,
					fmt.Sprintf("Message ID: %s, Body: %s", message.MessageId, message.Body)))
				processedMessages = append(processedMessages, makeEntry(message))
				continue
			}

			payload := sms.Payload{
				PhoneNumber: req.PhoneNumber,
				Message:     req.Message,
			}

			if err := gw.Send(payload); err != nil {
				result = multierror.Append(result, err)
			} else {
				processedMessages = append(processedMessages, makeEntry(message))
			}
		}

		if err := result.ErrorOrNil(); err != nil {
			s.DeleteMessageBatch(&sqs.DeleteMessageBatchInput{
				Entries:  processedMessages,
				QueueUrl: &sqsEvent.Records[0].EventSourceARN,
			})

			return err
		}

		return nil
	}, nil
}

func makeEntry(msg events.SQSMessage) *sqs.DeleteMessageBatchRequestEntry {
	return &sqs.DeleteMessageBatchRequestEntry{
		Id:            &msg.MessageId,
		ReceiptHandle: &msg.ReceiptHandle,
	}
}
