package Insta

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	database "RestAPI/database"
	uploadPost "RestAPI/fileHandler"
	getDat "RestAPI/user"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Article struct {
	id          string
	title       string
	description string
	urlToImage  string
	publishedAt time.Time
	fileName    string
	userID	  	string
}


func AddPostUser(userID string, fileName string) error {
	conn := database.InitiateMongoClient();
	collection := conn.Database("golangREST").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(ctx, 
		bson.M{"id": userID}, 
		bson.M{"$push": bson.M{"post": fileName}})

	return err
}


func AddInstaPost(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		r.ParseForm()
		post := Article{
			id:          	uuid.New().String(),
			title:       	r.Form["title"][0],
			description: 	r.Form["description"][0],
			urlToImage:  	r.Form["urlToImage"][0],
			publishedAt: 	time.Now(),
			fileName:    	uuid.New().String() + ".jpg",
			userID:			r.Form["userID"][0],
		}
		user := getDat.FetchUser(post.userID)

		if (user == nil) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(bson.M{"message": "No document with given UID"})
			return
		} else{
			
			err := AddPostUser(post.userID, post.fileName)
			if (err != nil) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(bson.M{"message": "Error updating user db"})
				return 
			}

			conn := database.InitiateMongoClient()
			db := conn.Database("golangREST")
			collection := db.Collection("InstPost")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			result, err := collection.InsertOne(ctx, bson.M{
				"id":          post.id,
				"title":       post.title,
				"description": post.description,
				"urlToImage":  post.urlToImage,
				"publishedAt": post.publishedAt,
				"fileName":    post.fileName,
			})
			filename := post.fileName
			fileSize := uploadPost.UploadFile(post.urlToImage, filename)
	
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(bson.M{"message": "Unable to upload document"})
			}else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(bson.M{
					"message": 	"Successfully uploaded",
					"id":      	result.InsertedID,
					"fileName": filename,
					"fileSize": fileSize,
					"user": 	user,
					"post": 	post,
				})
				log.Printf("Write Data Successfully")
			}
			
		}
	}
}
