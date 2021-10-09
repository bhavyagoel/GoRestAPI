package user

import (
	database "RestAPI/database"
	"context"
	"encoding/json"
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "User not found"})
			// fmt.Fprintf(w, "No Document with id: %s\n", id)
		}else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		    json.NewEncoder(w).Encode(user);
			// fmt.Fprintf(w, "User found with id: %s\n", id)
			// fmt.Fprintf(w, "User found with email: %s\n", user["Email"])
			// fmt.Fprintf(w, "User found with name: %v\n", user["name"])
			// fmt.Fprintf(w, "User found with hashed Password: %s\n", user["Password"])
			// fmt.Fprintf(w, "User found with posts: %v\n", user["Posts"])		
		}
	} else {
		w.WriteHeader((http.StatusMethodNotAllowed))
		fmt.Fprintf(w, "Invalid Request")
	}
}


