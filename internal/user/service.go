package user

import (
	"context"
	"log"

	"github.com/existing-test/internal/hostel"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var userColleciton *mongo.Collection
var bookingCollection *mongo.Collection

// UserCollection is the database connection for user colleciton
func UserCollection(c *mongo.Database) {
	userColleciton = c.Collection("user")
}

// BookingCollection is the database connection for booking collection
func BookingCollection(c *mongo.Database) {
	bookingCollection = c.Collection("booking")
}

func getUserDetail(userID string) (*UserWithoutPassword, error) {
	user := User{}
	err := userColleciton.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a user because of %v", err)
		return nil, err
	}

	return user.removeUserPassword(), nil
}

func bookHostel(userID string, hostelID string) error {
	booking := Booking{
		HostelID: hostelID,
		UserID:   userID,
	}
	err := hostel.BookHostel(hostelID)
	if err != nil {
		log.Printf("Error while booking a hostel, because of %v", err)
		return err
	}
	_, err = bookingCollection.InsertOne(context.TODO(), booking)
	if err != nil {
		log.Printf("Error while booking a hostel, because of %v", err)
		return err
	}

	return nil
}

func getBookings(userID string) ([]Booking, error) {
	bookings := []Booking{}
	cursor, err := bookingCollection.Find(context.TODO(), bson.M{"userId": userID})

	if err != nil {
		log.Printf("Error while getting user's bookings because of %v", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var booking Booking
		cursor.Decode(&booking)
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func getBookingDetail(bookingID string) (*Booking, error) {
	booking := Booking{}
	err := bookingCollection.FindOne(context.TODO(), bson.M{"_id": bookingID}).Decode(&booking)
	if err != nil {
		log.Printf("Error while getting a booking information.")
		return nil, err
	}
	return &booking, nil
}

// GetUserByUsername will query a user by username
func GetUserByUsername(username string) (*UserWithoutPassword, error) {
	user := User{}
	err := userColleciton.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a user by username.")
		return nil, err
	}
	return user.removeUserPassword(), nil
}
