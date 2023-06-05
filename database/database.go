/**
 * Created by VoidArtanis on 10/22/2017
 */

package shared

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var db *firestore.Client
var app *firebase.App
var err error

// Initializing the database
func InitDatabase() {
	ctx := context.Background()
	sa := option.WithCredentialsFile(".config/firebase_secret_key.json")
	app, err = firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	db, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer CloseDatabase()
}

// Return the instance of database
func GetDatabase() *firestore.Client {
	return db
}

// Close the databas
func CloseDatabase() {
	db.Close()
}
