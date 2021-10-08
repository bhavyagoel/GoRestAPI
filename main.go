package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type article struct {
	id 			string
	title 		string 
	description string 
	urlToImage 	string 
	publishedAt time.Time 
	fileName 	string 
}

func InitiateMongoClient() *mongo.Client {
    var err error
    var client *mongo.Client
    uri := "mongodb://localhost:27017"
    opts := options.Client()
    opts.ApplyURI(uri)
    opts.SetMaxPoolSize(5)
    if client, err = mongo.Connect(context.Background(), opts); err != nil {
        fmt.Println(err.Error())
    }
    return client
}

func addInstaPost(w http.ResponseWriter, r *http.Request) {
	conn := InitiateMongoClient()
	db := conn.Database("AppointyREST")
	collection := db.Collection("InstPost")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	r.ParseForm()
	post := article{
		id 			: uuid.New().String(),
		title 		: r.Form["title"][0],
		description : r.Form["description"][0],
		urlToImage	: r.Form["urlToImage"][0],
		publishedAt : time.Now(),
		fileName	: uuid.New().String() + ".jpg",
	}

	// result, err := collection.InsertOne(ctx, post)
	result, err := collection.InsertOne(ctx, bson.M{
		"id" 			: post.id,
		"title" 		: post.title,
		"description" 	: post.description,
		"urlToImage" 	: post.urlToImage,
		"publishedAt" 	: post.publishedAt,
		"fileName" 		: post.fileName,
	})

	filename := post.fileName
    fileSize := UploadFile(post.urlToImage, filename)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Inserted a single document: %v\n", result.InsertedID)
	fmt.Fprintf(w, "File size: %v\n", fileSize)
	log.Printf("Write Data Successfully")
}

func UploadFile(file, filename string, ) int {

    data, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    conn := InitiateMongoClient()
    bucket, err := gridfs.NewBucket(
        conn.Database("AppointyREST"),
    )
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    uploadStream, err := bucket.OpenUploadStream(
        filename,
    )
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer uploadStream.Close()

    fileSize, err := uploadStream.Write(data)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
	return fileSize
}

func DownloadFile(fileName string) {
    conn := InitiateMongoClient()

    // For CRUD operations, here is an example
    db := conn.Database("AppointyREST")
    fsFiles := db.Collection("fs.files")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    var results bson.M
    err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
    if err != nil {
        log.Fatal(err)
    }
    // you can print out the results
    fmt.Println(results)

    bucket, _ := gridfs.NewBucket(
        db,
    )
    var buf bytes.Buffer
    dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("File size to download: %v\n", dStream)
    ioutil.WriteFile(fileName, buf.Bytes(), 0600)

}

func main() {
    // Uncomment the below line and comment the UploadFile above this line to download the file
    DownloadFile("cc5d692d-c44a-4baf-8cdb-7203280eaa7d.jpg")

	http.HandleFunc("/addInstaPost", addInstaPost)
	http.ListenAndServe(":8080", nil)
	
}
