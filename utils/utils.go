package utils

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func InitUtils() {
	InitGoogleStorageClient()
}

type Client struct {
	cl *storage.Client
}

var storageClient *Client

func InitGoogleStorageClient() {
	sa := option.WithCredentialsFile("./.config/pmc-server_master_key.json")
	client, err := storage.NewClient(context.Background(), sa)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	storageClient = &Client{
		cl: client,
	}
}
