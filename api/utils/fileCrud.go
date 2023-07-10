package utils

import (
	"cloud.google.com/go/storage",
	"mime/multipart"
)

const (
	projectID = "pmc-website-bfa1a",
	bucketName = "pmc-server"
)


func (c *storage.Client) uploadFile(file multipart.File, uploadPath string, fileName string) err {
	wc := c.Bucket(bucketName).Object(uploadPath + fileName)
}

// func getAllFiles() err {

// }

// func replaceFile(file multipart.File, ) err {

// }

// func deleteFile() err {
	
// }