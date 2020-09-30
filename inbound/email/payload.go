package email

type jsonPayload struct {
	ReceiverEmails []string `json:"receivers"`
	SenderEmail    string   `json:"sender"`
	Subject        string   `json:"subject"`
	HTML           string   `json:"html"`
	Text           string   `json:"text"`
}
