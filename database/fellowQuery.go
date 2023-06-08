package database

import (
	"context"

	"github.com/UBC-Product-Management-Club/website-server/models"
	"google.golang.org/api/iterator"
)

// Add a greeting to database
func AddFellow(fellow models.Fellow) error {
	_, _, err := client.Collection("fellow").Add(context.Background(), fellow)
	return err
}

// Get all greetings from database
func GetAllFellows() ([]models.Fellow, error) {
	var fellows []models.Fellow
	iter := client.Collection("fellow").Documents(context.Background())
	for doc, err := iter.Next(); err != iterator.Done; doc, err = iter.Next() {
		if err != nil {
			return []models.Fellow{}, err
		}

		var tempFellow models.Fellow
		if err := doc.DataTo(&tempFellow); err != nil {
			return []models.Fellow{}, err
		}
		fellows = append(fellows, tempFellow)
	}
	return fellows, nil
}
