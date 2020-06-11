package controllers

import (
	"context"
	"math/rand"

	"appstud.com/github-core/src/db"
	"appstud.com/github-core/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// this func generate a random string with a ${n} lenght
func generateToken(n int, c *gin.Context) string {
	usersCollection := db.GetUserCollection()
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	//Check if the token is already assigned
	count, err := usersCollection.
		CountDocuments(context.TODO(), bson.D{primitive.E{Key: "token", Value: string(b)}})
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	if count > 0 {
		return generateToken(256, c)
	}

	return string(b)
}

// This method allow to login a user (api/users/register)
func registerUser(c *gin.Context) {
	usersCollection := db.GetUserCollection()
	var generatedToken = generateToken(256, c)

	//check if params are valid
	if len(c.Query("username")) == 0 || len(c.Query("password")) == 0 {
		c.JSON(403, "Username/password cannot be empty")
		panic("Username/password cannot be empty")
	}

	//check if the username is already taken
	count, err := usersCollection.
		CountDocuments(context.TODO(), bson.D{primitive.E{Key: "username", Value: c.Query("username")}})

	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}

	if count == 0 {
		_, err := usersCollection.InsertOne(context.TODO(), bson.D{
			{Key: "username", Value: c.Query("username")},
			{Key: "password", Value: c.Query("password")},
			{Key: "token", Value: generatedToken},
		})
		if err != nil {
			c.JSON(500, "Internal Server Error")
			panic(err)
		}
		c.JSON(200, models.AuthentificationResponse{
			Username: c.Query("username"),
			Token:    generatedToken,
		})
	} else {
		c.JSON(409, "Username is already taken")
	}
}

// This method return all the registered users (api/users/)
func getAllUsers(c *gin.Context) {
	usersCollection := db.GetUserCollection()
	allUserName := models.UsersResponses{}
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}
	err = cursor.All(context.TODO(), &allUserName)
	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}

	c.JSON(200, allUserName)
}

// This method allow to login a user (api/users/login)
func loginUser(c *gin.Context) {
	usersCollection := db.GetUserCollection()
	currentUser := models.User{}
	var generatedToken = generateToken(256, c)

	if len(c.Query("username")) == 0 || len(c.Query("password")) == 0 {
		c.JSON(403, "Username/password cannot be empty")
		panic("Username/password cannot be empty")
	}
	err := usersCollection.FindOne(context.TODO(), bson.M{"username": c.Query("username")}).Decode(&currentUser)

	if err != nil {
		c.JSON(500, "Internal Server Error")
		panic(err)
	}

	if currentUser.Password != c.Query("password") {
		c.JSON(403, "Invalid password")
	} else {
		currentUser.Token = generatedToken
		_, err := usersCollection.
			UpdateOne(context.TODO(), bson.M{"username": c.Query("username")},
				bson.D{{"$set",
					bson.D{
						{"token", generatedToken},
					},
				}})

		if err != nil {
			c.JSON(500, "Internal Server Error")
			panic(err)
		}

		c.JSON(200, models.AuthentificationResponse{Username: currentUser.Username, Token: currentUser.Token})
	}
}

// This method allow you to get your username by sending your token (api/users/me)
func getUserByToken(c *gin.Context) {
	usersCollection := db.GetUserCollection()
	currentUser := models.UserName{}
	err := usersCollection.FindOne(context.TODO(), bson.M{"token": c.Query("token")}).Decode(&currentUser)

	if err != nil {
		c.JSON(403, "Invalid token!")
		panic(err)
	}
	c.JSON(200, currentUser)

}

// AuthentificationController - Route controller
func AuthentificationController(engine *gin.Engine) {
	engine.GET("/api/users/login/", loginUser)
	engine.GET("/api/users/me/", getUserByToken)
	engine.GET("/api/users/register/", registerUser)
	engine.GET("/api/users/", getAllUsers)

}
