package ifttt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/finfech/notiplat-gateway/domain/useCase/sms"
)

type iftttSMSGW struct {
	key string
}

type payload struct {
	value1 string `json:"value1"`
}

// NewIftttSMSGateway godoc
func NewIftttSMSGateway() (sms.Gateway, error) {
	key := os.Getenv("IFTTT_KEY")
	if key == "" {
		return nil, errors.New("IFTTT_KEY is not defined...")
	}

	return &iftttSMSGW{
		key: key,
	}, nil
}

func (gw *iftttSMSGW) Send(p sms.Payload) error {
	values := map[string]string{
		"value1": p.Message,
	}

	reqBody, err := json.Marshal(values)
	if err != nil {
		return err
	}

	url := "https://maker.ifttt.com/trigger/sms/with/key/" + gw.key
	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	fmt.Println(res, err)
	return err
}
