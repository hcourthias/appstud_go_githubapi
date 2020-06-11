package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client
var usersCollection *mongo.Collection

// Init the mongoDB
func Init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://appstud:1234567890@cluster0-vdw34.mongodb.net/github-api?retryWrites=true&w=majority")
	db, err := mongo.Connect(context.TODO(), clientOptions) // connect to mongoDB
	if err != nil {
		panic(err)
	}
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	usersCollection = db.Database("github-api").Collection("users")
}

// GetUserCollection allow to get the users collection
func GetUserCollection() *mongo.Collection {
	return usersCollection
}
