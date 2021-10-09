package fileHandler

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"time"

	database "RestAPI/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)


func DownloadFile(fileName string) int64 {
	conn := database.InitiateMongoClient()

	// For CRUD operations, here is an example
	db := conn.Database("AppointyREST")
	fsFiles := db.Collection("fs.files")
	ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	var results bson.M
	err := fsFiles.FindOne(ctx, bson.M{}).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}

	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(fileName, &buf)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(fileName, buf.Bytes(), 0600)
	return dStream
}