package hostel

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func HostelCollection(c *mongo.Database) {
	collection = c.Collection("hostel")
}

func getHostels() ([]Hostel, error) {
	hostels := []Hostel{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all hostels because of %v", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var hostel Hostel
		cursor.Decode(&hostel)
		hostels = append(hostels, hostel)
	}

	return hostels, nil
}

func GetHostelById(hostelId string) (*Hostel, error) {
	hostel := Hostel{}
	err := collection.FindOne(context.TODO(), bson.M{"id": hostelId}).Decode(&hostel)
	if err != nil {
		log.Printf("Error while getting a hostel becase of %v", err)
		return nil, err
	}

	return &hostel, nil
}

func CreateHostel(hostel Hostel) error {

	name := hostel.Name

	newHostel := Hostel{
		ID:   uuid.New().String(),
		Name: name,
	}

	_, err := collection.InsertOne(context.TODO(), newHostel)

	if err != nil {
		log.Printf("Error while inserting hostel because of %v", err)
		return err
	}

	return nil
}

func CheckIfHostelExists(name string) bool {
	var existed Hostel
	doc := collection.FindOne(context.TODO(), bson.M{"name": name})
	doc.Decode(&existed)

	if existed != (Hostel{}) {

		return true
	}
	return false
}
