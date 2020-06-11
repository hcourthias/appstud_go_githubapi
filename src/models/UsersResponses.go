package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `bson:"password" json:"password"`
	Token    string             `json:"token" bson:"token"`
}

// UserName : I created this type because I struggled on making a select on the mongo.Find :(
type UserName struct {
	Username string `json:"username" bson:"username"`
}

// UsersResponses basic /api response
type UsersResponses []UserName
