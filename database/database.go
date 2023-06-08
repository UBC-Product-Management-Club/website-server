/**
 * Created by VoidArtanis on 10/22/2017
 */

package database

import (
	"context"

	"cloud.google.com/go/firestore"
)

var client *firestore.Client

// Initialize the database
func InitDatabase() error {
	var err error
	projectID := "pmc-website-bfa1a"
	client, err = firestore.NewClient(context.Background(), projectID)
	return err
}

// Return the instance of database
func GetDatabase() *firestore.Client {
	return client
}

// Close the databas
func CloseDatabase() {
	client.Close()
}
