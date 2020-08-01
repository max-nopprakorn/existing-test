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

func BookingCollection(c *mongo.Database) {
	bookingCollection = c.Collection("booking")
}

func getUserDetail(userID string) (*User, error) {
	user := User{}
	err := userColleciton.FindOne(context.TODO(), bson.M{"id": userID}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a user because of %v", err)
		return nil, err
	}

	return &user, nil
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

func getBookingDetail(bookingId string) (*Booking, error) {
	booking := Booking{}
	err := bookingCollection.FindOne(context.TODO(), bson.M{"id": bookingId}).Decode(&booking)
	if err != nil {
		log.Printf("Error while getting a booking information.")
		return nil, err
	}
	return &booking, err
}
