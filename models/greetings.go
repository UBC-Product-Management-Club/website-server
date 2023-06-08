package models

import "fmt"

type Greeting struct {
	Message string `json:"message" binding:"required" firestore:"message"`
}

func (this Greeting) Greet() (string, error) {
	if this.Message == "" {
		return "", fmt.Errorf("ERROR: empty greeting")
	}
	return this.Message, nil
}
