package user

import (
	auth "RestAPI/auth"
	database "RestAPI/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type users struct {
	id 			string 
	name		string
	Email 		string
	Password 	string
	Posts		*[]string
}

func addUser(w http.ResponseWriter, r *http.Request) {

	if(r.Method == "POST") {
		conn := database.InitiateMongoClient()
		db := conn.Database("golangREST")
		collection := db.Collection("Users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() 

		r.ParseForm()
		user := users{
			id: 		uuid.New().String(),
			name: 		r.Form["name"][0],
			Email: 		r.Form["email"][0],
			Password: 	r.Form["password"][0],
		}
		hash, err := auth.HashPassword(user.Password)
		if (err!=nil) {
			log.Fatal(err)
		}
		result, err := collection.InsertOne(ctx, bson.M{
			"id": 		user.id,
			"name": 	user.name,
			"Email": 	user.Email,
			"Password": hash, 
			"Posts": 	user.Posts,
		})
		if (err!=nil) {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "Inserted a single user document: %v\n", result.InsertedID)
		log.Printf("Write Data Successfully")
	}else {
		fmt.Fprintf(w, "Invalid Request")
	}
}

func init() {
	http.HandleFunc("/addUser", addUser)
}