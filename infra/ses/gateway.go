package ses

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/finfech/notiplat-gateway/domain/useCase/email"
)

type sesEmailGW struct {
	ses *ses.SES
}

// NewSESEmailGateway godoc
func NewSESEmailGateway(sess *session.Session) (email.Gateway, error) {
	return &sesEmailGW{
		ses: ses.New(sess),
	}, nil
}

const charSet = "UTF-8"

func (gw *sesEmailGW) Send(payload email.Payload) error {
	input := makeSESEmailInput(payload)

	_, err := gw.ses.SendEmail(input)
	if err != nil {
		return err
	}

	return nil
}

func makeSESEmailInput(payload email.Payload) *ses.SendEmailInput {
	toAddresses := []*string{}
	for _, addr := range payload.ReceiverEmails {
		toAddresses = append(toAddresses, aws.String(addr))
	}
	fmt.Println("payload", payload, toAddresses)
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(payload.HTML),
				},
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(payload.Text),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(payload.Subject),
			},
		},
		Source: aws.String(payload.SenderEmail),
	}
}
