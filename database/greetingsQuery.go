package database

import (
	"context"

	"github.com/UBC-Product-Management-Club/website-server/models"
	"google.golang.org/api/iterator"
)

// Add a greeting to database
func AddGreeting(msg models.Greeting) error {
	_, _, err := client.Collection("messages").Add(context.Background(), msg)
	return err
}

// Get all greetings from database
func GetAllGreetings() ([]models.Greeting, error) {
	var greetings []models.Greeting
	iter := client.Collection("messages").Documents(context.Background())
	for doc, err := iter.Next(); err != iterator.Done; doc, err = iter.Next() {
		if err != nil {
			return []models.Greeting{}, err
		}

		var tempGreet models.Greeting
		if err := doc.DataTo(&tempGreet); err != nil {
			return []models.Greeting{}, err
		}
		greetings = append(greetings, tempGreet)
	}
	return greetings, nil
}
