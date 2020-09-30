package sns

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/finfech/notiplat-gateway/domain/useCase/sms"
)

type snsSmsGW struct {
	sns *sns.SNS
}

// NewSNSSmsGateway godoc
func NewSNSSmsGateway(sess *session.Session) (sms.Gateway, error) {
	return &snsSmsGW{
		sns: sns.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")}),
	}, nil
}

func (gw *snsSmsGW) Send(p sms.Payload) error {
	if p.PhoneNumber == "" {
		fmt.Println("Ignore empty phone number")
		return nil
	}

	input := &sns.PublishInput{
		Message:          aws.String(p.Message),
		MessageStructure: aws.String("text"),
		PhoneNumber:      aws.String(p.PhoneNumber),
	}

	if _, err := gw.sns.Publish(input); err != nil {
		return err
	}

	return nil
}
