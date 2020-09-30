package email

// Payload godoc
type Payload struct {
	ReceiverEmails []string
	SenderEmail    string
	Subject        string
	HTML           string
	Text           string
}
