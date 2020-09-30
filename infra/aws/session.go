package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession godoc
func NewSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
}
