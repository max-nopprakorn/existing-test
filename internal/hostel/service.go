package hostel

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Collection is the database connection for hostel colleciton
func Collection(c *mongo.Database) {
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

func getAvailableHostels() ([]Hostel, error) {
	availableHostels := []Hostel{}
	cursor, err := collection.Find(context.TODO(), bson.M{"isAvailable": true})

	if err != nil {
		log.Printf("Error while getting all hostels because of %v", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var hostel Hostel
		cursor.Decode(&hostel)
		availableHostels = append(availableHostels, hostel)
	}

	return availableHostels, nil
}

func getHostelByID(hostelID string) (*Hostel, error) {
	hostel := Hostel{}
	err := collection.FindOne(context.TODO(), bson.M{"id": hostelID}).Decode(&hostel)
	if err != nil {
		log.Printf("Error while getting a hostel becase of %v", err)
		return nil, err
	}

	return &hostel, nil
}

func createHostel(hostel Hostel) error {

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

func checkIfDuplicate(name string) bool {
	var existed Hostel
	collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&existed)

	if existed != (Hostel{}) {
		return true
	}
	return false
}

// BookHostel will set the hostel to unavailable
func BookHostel(hostelID string) error {
	booked := bson.M{
		"$set": bson.M{
			"isAvaliable": false,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"id": hostelID}, booked)

	if err != nil {
		log.Printf("Error while booking a hostel because of %v", err)
		return err
	}

	return nil
}

// CheckIfAvaliable will check if hostel is available or not
func CheckIfAvaliable(hostelID string) bool {
	var hostel Hostel
	collection.FindOne(context.TODO(), bson.M{"id": hostelID}).Decode(&hostel)

	if hostel != (Hostel{}) && hostel.IsAvailable {
		return true
	}
	return false
}
