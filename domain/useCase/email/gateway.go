package email

// Gateway godoc
type Gateway interface {
	Send(payload Payload) error
}
