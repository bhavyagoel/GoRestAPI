package user

import (
	auth "RestAPI/auth"
	database "RestAPI/database"
	"context"
	"encoding/json"
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
	post		*[]string
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
			usrName: 	r.Form["name"][0],
			Email: 		r.Form["email"][0],
			Password: 	r.Form["password"][0],
		}
		hash, err := auth.HashPassword(user.Password)
		if (err!=nil) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "Unable to hash password"})
			return 
		}
		result, err := collection.InsertOne(ctx, bson.M{
			"id": 		user.id,
			"name": 	user.usrName,
			"Email": 	user.Email,
			"Password": hash, 
			"post": 	user.post,
		})
		if (err!=nil) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "Unable to insert database"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{
			"message" : "User added successfully",
			"result" : result,
		});

		log.Printf("Write Data Successfully")
	}else {
		w.WriteHeader((http.StatusMethodNotAllowed))
		fmt.Fprintf(w, "Invalid Request")
	}
}
