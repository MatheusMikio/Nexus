package models

type ErrorMessage struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

func NewErrorMessage(property, message string) *ErrorMessage {
	return &ErrorMessage{
		Property: property,
		Message:  message,
	}
}
