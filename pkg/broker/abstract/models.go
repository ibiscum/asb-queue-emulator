package abstract

type Message struct {
	MessageId string `json:"messageId"`
	Body      []byte `json:"body"`
}
