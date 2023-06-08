package database

import (
	"context"

	"github.com/UBC-Product-Management-Club/website-server/models"
	"google.golang.org/api/iterator"
)

// Add an executive to database
func AddExecutive(exec models.Executive) error {
	_, _, err := client.Collection("executive").Add(context.Background(), exec)
	return err
}

// Get all executives from database
func GetAllExecutives() ([]models.Executive, error) {
	var execs []models.Executive
	iter := client.Collection("executive").Documents(context.Background())
	for doc, err := iter.Next(); err != iterator.Done; doc, err = iter.Next() {
		if err != nil {
			return []models.Executive{}, err
		}

		var tempExec models.Executive
		if err := doc.DataTo(&tempExec); err != nil {
			return []models.Executive{}, err
		}
		execs = append(execs, tempExec)
	}
	return execs, nil
}
