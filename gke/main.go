package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/storage"

	"cloud.google.com/go/compute/metadata"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
)

var ()

const (
	bucketName = "some-bucket"
)

func gethandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	if !metadata.OnGCE() {
		log.Printf("app not running on compute")
		http.Error(w, "app not running on compute", http.StatusBadRequest)
	}

	object := "file1.txt"
	expires := time.Now().Add(time.Minute * 10)

	// to use the default metadata service account
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating storage client: %v", err)
		http.Error(w, "Error creating storage client", http.StatusBadRequest)
	}
	s, err := storageClient.Bucket(bucketName).SignedURL(object, &storage.SignedURLOptions{
		Method:  http.MethodGet,
		Expires: expires,
		Scheme:  storage.SigningSchemeV4,
	})

	// to use a different service account, the default sa must have tokencreator role on the targetPrincipal

	// targetPrincipal := "othersigner@project.iam.gserviceaccount.com"

	// delegates := []string{}
	// impersonatedTS, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
	// 	TargetPrincipal: targetPrincipal,
	// 	Scopes:          []string{"https://www.googleapis.com/auth/cloud-platform"},
	// 	Delegates:       delegates,
	// })
	// if err != nil {
	// 	log.Printf("Error impersonating target SA: %v", err)
	// 	http.Error(w, "Error impersonating target SA", http.StatusBadRequest)
	// }

	// storageClient, err := storage.NewClient(ctx, option.WithTokenSource(impersonatedTS))
	// if err != nil {
	// 	log.Printf("Error creating storage client: %v", err)
	// 	http.Error(w, "Error creating storage client", http.StatusBadRequest)
	// }

	// s, err := storageClient.Bucket(bucketName).SignedURL(object, &storage.SignedURLOptions{
	// 	GoogleAccessID: targetPrincipal,
	// 	Method:         http.MethodGet,
	// 	Expires:        expires,
	// 	Scheme:         storage.SigningSchemeV4,
	// })

	if err != nil {
		log.Printf("Error creating signedurl: %v", err)
		http.Error(w, "Error creating signedurl:", http.StatusBadRequest)
	}

	log.Println(s)
	fmt.Fprint(w, fmt.Sprintf("%s", s))
}

func main() {

	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/").HandlerFunc(gethandler)
	var server *http.Server
	server = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	http2.ConfigureServer(server, &http2.Server{})
	fmt.Println("Starting Server..")
	err := server.ListenAndServe()
	fmt.Printf("Unable to start Server %v", err)

}
