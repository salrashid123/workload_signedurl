package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
)

var ()

const (
	targetPrincipal = "urlsigner@fabled-ray-104117.iam.gserviceaccount.com"
	bucketName      = "some-bucket"
)

func main() {
	ctx := context.Background()
	object := "file1.txt"
	expires := time.Now().Add(time.Minute * 10)

	// https://pkg.go.dev/cloud.google.com/go/storage#hdr-Credential_requirements_for_signing
	storageClient, err := storage.NewClient(ctx)
	s, err := storageClient.Bucket(bucketName).SignedURL(object, &storage.SignedURLOptions{
		GoogleAccessID: targetPrincipal,
		Method:         http.MethodGet,
		Expires:        expires,
		Scheme:         storage.SigningSchemeV4,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)

}
