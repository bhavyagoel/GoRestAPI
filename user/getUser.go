package user

import (
	database "RestAPI/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func FetchUser(userID string) bson.M {
	conn := database.InitiateMongoClient();
	collection := conn.Database("golangREST").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user bson.M
	err := collection.FindOne(ctx, bson.M{"id": userID}).Decode(&user)
	if (err != nil) {
		return nil
	} 
	return user
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	if(r.Method=="GET") {

		r.ParseForm()
		id := r.Form["id"][0]

		user := FetchUser(id)
		if (user == nil) {
			fmt.Fprintf(w, "No Document with id: %s\n", id)
		}else {
			fmt.Fprintf(w, "User found with id: %s\n", id)
			fmt.Fprintf(w, "User found with email: %s\n", user["Email"])
			fmt.Fprintf(w, "User found with name: %v\n", user["name"])
			fmt.Fprintf(w, "User found with hashed Password: %s\n", user["Password"])
			fmt.Fprintf(w, "User found with posts: %v\n", user["Posts"])		
		}
	} else {
		fmt.Fprintf(w, "Invalid Request")
	}
}


