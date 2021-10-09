package Insta

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	database "RestAPI/database"
	uploadPost "RestAPI/fileHandler"
	getDat "RestAPI/user"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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


func AddPostUser(userID string, fileName string) (mongo.UpdateResult, error) {
	conn := database.InitiateMongoClient();
	collection := conn.Database("golangREST").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, 
		bson.M{"id": userID}, 
		bson.M{"$push": bson.M{"post": fileName}})
	return *result, err
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
			fmt.Fprintf(w, "No Document with id: %s\n", post.userID)

		} else{
			
			res, err := AddPostUser(post.userID, post.fileName)
			if (err != nil) {
				fmt.Fprintf(w, "Error updating user db: %s\n", err)
				panic(err)
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
				fmt.Fprintf(w, "Unable to upload document")
			}else {
				fmt.Fprintf(w, "Inserted a single document: %v\n", result.InsertedID)
				fmt.Fprintf(w, "Updated user with ID: %v\n", res.UpsertedID)
				fmt.Fprintf(w, "File size: %v\n", fileSize)
				log.Printf("Write Data Successfully")
			}
			
		}
	}
}
