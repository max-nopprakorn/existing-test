package hostel

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func HostelCollection(c *mongo.Database) {
	collection = c.Collection("hostel")
}

func GetHostels(c *gin.Context) {
	hostels := []Hostel{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all hostels because of %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong.",
		})
	}

	for cursor.Next(context.TODO()) {
		var hostel Hostel
		cursor.Decode(&hostel)
		hostels = append(hostels, hostel)
	}

	c.JSON(http.StatusOK, hostels)
}

func GetHostelById(c *gin.Context) {
	hostelId := c.Param("hostelID")
	hostel := Hostel{}
	err := collection.FindOne(context.TODO(), bson.M{"id": hostelId}).Decode(&hostel)
	if err != nil {
		log.Printf("Error while getting a hostel becase of %v", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Resource not found",
		})
		return
	}

	c.JSON(200, hostel)
}

func CreateHostel(c *gin.Context) {
	var hostel Hostel
	err := c.ShouldBindJSON(&hostel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	name := hostel.Name

	newHostel := Hostel{
		ID:   uuid.New().String(),
		Name: name,
	}

	var existed Hostel
	doc := collection.FindOne(context.TODO(), bson.M{"name": name})
	doc.Decode(&existed)

	if existed != (Hostel{}) {
		c.JSON(409, gin.H{
			"message": "Name already exists.",
		})
		return
	}

	_, err = collection.InsertOne(context.TODO(), newHostel)

	if err != nil {
		log.Printf("Error while inserting hostel because of %v", err)
		c.JSON(500, gin.H{
			"message": "somethign went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, newHostel)
	return
}
