package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

type User struct {
	id       string
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Setting Up MongoDB Connection
// func init() {
	//  defer cancel()
	//     var err error
	// fmt.Println("HELLOOOOOO")
	//  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	//     if err != nil {
	//         log.Fatal(err)
	//     }
// }

func addUser(w http.ResponseWriter, r *http.Request) {
    var err error
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	r.ParseForm()
	collection := client.Database("appointy").Collection("users")
	u1 := User{
		id:       uuid.New().String(),
		Email:    r.Form["email"][0],
		Password: r.Form["password"][0],
	}
	// json_data, err := json.Marshal(u1)
	// if(err != nil){
	//        panic(err)
	// }
	fmt.Println(u1)
	fmt.Println("\n break Line")
	res, err := collection.InsertOne(ctx, bson.M{
		"id":       u1.id,
		"email":    u1.Email,
		"password": u1.Password,
	})
	if err != nil {
		panic(err)
	}
	id := res.InsertedID
	fmt.Println(id)
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
	http.HandleFunc("/", sayhelloName) // set router
	http.HandleFunc("/login", addUser)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
