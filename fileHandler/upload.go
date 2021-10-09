package fileHandler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	database "RestAPI/database"

	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func UploadFile(file, filename string) int {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	conn := database.InitiateMongoClient()
	bucket, err := gridfs.NewBucket(
		conn.Database("golangREST"),
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
