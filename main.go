package main

import (
	insta "RestAPI/Insta"
	user "RestAPI/user"
	"net/http"
)

func main() {
	// Uncomment the below line and comment the UploadFile above this line to download the file
	// DownloadFile("cc5d692d-c44a-4baf-8cdb-7203280eaa7d.jpg")
	
	http.HandleFunc("/addInstaPost", insta.AddInstaPost)
	http.HandleFunc("/user/getUserById", user.GetUserById)
	http.HandleFunc("/user/addUser", user.AddUser)

	http.ListenAndServe(":8080", nil)
}
