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

type Users struct {
	id 			string 
	usrName		string
	Email 		string
	Password 	string
	Posts		*[]string
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if(r.Method == "POST") {
		conn := database.InitiateMongoClient()
		db := conn.Database("golangREST")
		collection := db.Collection("Users")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() 

		r.ParseForm()
		user := Users{
			id: 		uuid.New().String(),
			usrName: 		r.Form["name"][0],
			Email: 		r.Form["email"][0],
			Password: 	r.Form["password"][0],
		}
		hash, err := auth.HashPassword(user.Password)
		if (err!=nil) {
			fmt.Fprintf(w, "Unable to hash password. \n%s", err)
		}
		result, err := collection.InsertOne(ctx, bson.M{
			"id": 		user.id,
			"name": 	user.usrName,
			"Email": 	user.Email,
			"Password": hash, 
			"Posts": 	user.Posts,
		})
		if (err!=nil) {
			fmt.Fprintf(w, "Unable to insert database. \n%s", err)
		}

		fmt.Fprintf(w, "Inserted a single user document: %v\n", result.InsertedID)
		fmt.Fprintf(w, "User UUID: %v\n", user.id)

		log.Printf("Write Data Successfully")
	}else {
		fmt.Fprintf(w, "Invalid Request")
	}
}
