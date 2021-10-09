package Insta

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	database "RestAPI/database"
	uploadPost "RestAPI/fileHandler"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type article struct {
	id          string
	title       string
	description string
	urlToImage  string
	publishedAt time.Time
	fileName    string
}

func addInstaPost(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		conn := database.InitiateMongoClient()
		db := conn.Database("AppointyREST")
		collection := db.Collection("InstPost")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		r.ParseForm()
		post := article{
			id:          uuid.New().String(),
			title:       r.Form["title"][0],
			description: r.Form["description"][0],
			urlToImage:  r.Form["urlToImage"][0],
			publishedAt: time.Now(),
			fileName:    uuid.New().String() + ".jpg",
		}

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
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Inserted a single document: %v\n", result.InsertedID)
		fmt.Fprintf(w, "File size: %v\n", fileSize)
		log.Printf("Write Data Successfully")
	} else {
		fmt.Fprintf(w, "Invalid Request")
	}	
}

func init() {
	http.HandleFunc("/addInstaPost", addInstaPost)
}