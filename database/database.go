/**
 * Created by VoidArtanis on 10/22/2017
 */

package database

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var client *firestore.Client

// Initialize the database
func InitDatabase() {
	var err error
	projectID := "pmc-website-bfa1a"
	sa := option.WithCredentialsFile("./.config/firebase_secret_key.json")
	client, err = firestore.NewClient(context.Background(), projectID, sa)
	if err != nil {
		panic(err)
	}
}

// Return the instance of database
func GetDatabase() *firestore.Client {
	return client
}

// Close the databas
func CloseDatabase() {
	client.Close()
}
