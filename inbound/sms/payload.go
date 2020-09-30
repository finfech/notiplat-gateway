package sms

type jsonPayload struct {
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
}
