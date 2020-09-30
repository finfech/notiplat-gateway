package sms

// Gateway godoc
type Gateway interface {
	Send(payload Payload) error
}
