package auth

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/existing-test/helper"
	userHelper "github.com/existing-test/internal/user"
)

var userCollection *mongo.Collection

// UserCollection is the database connecion for user collection
func UserCollection(c *mongo.Database) {
	userCollection = c.Collection("user")
}

func login(req LoginRequest) (string, error) {
	user, _ := userHelper.GetUserByUsername(req.Username)
	token, err := helper.CreateToken(user.ID.Hex())
	if err != nil {
		log.Printf("Error while creating a token.")
		return "", err
	}
	return token, nil
}

func comparePassword(loginFrom LoginRequest) bool {
	user := userHelper.User{}
	userCollection.FindOne(context.TODO(), bson.M{"username": loginFrom.Username}).Decode(&user)
	if user == (userHelper.User{}) {
		return false
	}
	bytePassword := []byte(user.Password)
	byteRequestPassword := []byte(loginFrom.Password)

	err := bcrypt.CompareHashAndPassword(bytePassword, byteRequestPassword)

	if err != nil {
		return false
	}

	return true
}

func checkDuplicateUser(username string) bool {
	user := userHelper.User{}
	userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if user != (userHelper.User{}) {
		return true
	}
	return false
}

func register(req userHelper.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		log.Printf("Error while hasing a password, because of %v", err)
		return err
	}
	username := req.Username
	email := req.Email
	name := req.Name
	lastname := req.Lastname
	newUser := userHelper.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Email:    email,
		Name:     name,
		Lastname: lastname,
		Password: string(hashedPassword),
	}
	_, err = userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Printf("Error while creating a new user, becuase of %v", err)
		return err
	}
	return nil
}
