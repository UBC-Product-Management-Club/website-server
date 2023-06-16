package models

import (
	"fmt"
	"time"
)

// Error on the server
type ServerError struct {
	Metadata     metadata `json:"metadata"`
	ErrorMessage string   `json:"error_message"`
	StatusCode   int      `json:"-"`
}

// Metadata for the server's error
type metadata struct {
	Time        string `json:"time"`
	RequestType string `json:"request_type"`
}

// Create a new error
func NewServerError(msg, requestType string, statusCode int) ServerError {
	return ServerError{
		Metadata:     addMetadata(requestType),
		ErrorMessage: msg,
		StatusCode:   statusCode,
	}
}

// Return error logging message
func (se ServerError) Error() string {
	return fmt.Sprintf("[%s] %s [%v]", se.Metadata.Time, se.ErrorMessage, se.StatusCode)
}

// Return status code of the error
func (se ServerError) GetCode() int {
	return se.StatusCode
}

// Add metadata
func addMetadata(requestType string) metadata {
	return metadata{
		Time:        time.Now().Format(time.RFC3339),
		RequestType: requestType,
	}
}
