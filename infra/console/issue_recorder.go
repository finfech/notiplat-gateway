package console

import (
	"fmt"

	"github.com/finfech/notiplat-gateway/domain/useCase/issue"
)

type consoleIssueRecorder struct{}

// NewIssueRecorder godoc
func NewIssueRecorder() issue.Recorder {
	return &consoleIssueRecorder{}
}

func (r *consoleIssueRecorder) Record(err error) error {
	fmt.Println("Issued ", err)
	return nil
}
