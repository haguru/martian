package interfaces

type SMSClient interface {
	Send(outgoingDestination string, message string) (map[string]interface{},error)
	Receive(incomingDestination string, message string) (map[string]interface{}, error)
	// TODO: rethink this part as other sms api may not require this
	ValidIncomingMessage(apiKey string,timestamp string,requestSignature string, requestPayload string) bool
}