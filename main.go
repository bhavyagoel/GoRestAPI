package main

import (
	_ "RestAPI/Insta"
	_ "RestAPI/user"
	"net/http"
)

func main() {
	// Uncomment the below line and comment the UploadFile above this line to download the file
	// DownloadFile("cc5d692d-c44a-4baf-8cdb-7203280eaa7d.jpg")
	http.ListenAndServe(":8080", nil)
}
