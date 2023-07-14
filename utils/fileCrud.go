package utils

import (
	"context"
	"io"
	"mime/multipart"
	"time"
)

const (
	bucketName = "pmc-server-img"
)

func UploadFile(file multipart.File, uploadPath string, fileName string) (string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	wc := storageClient.cl.Bucket(bucketName).Object(uploadPath + fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return "https://storage.googleapis.com/" + bucketName + "/" + uploadPath + "/" + fileName, nil
}

// func getAllFiles() err {

// }

// func replaceFile(file multipart.File, ) err {

// }

// func deleteFile() err {

// }
