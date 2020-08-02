package user

import (
	"context"
	"log"

	"github.com/existing-test/internal/hostel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Printf("Error while transfrom string to objectId")
		return nil, err
	}
	err = userColleciton.FindOne(context.TODO(), bson.M{"_id": userObjectID}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a user because of %v", err)
		return nil, err
	}

	return user.removeUserPassword(), nil
}

func bookHostel(userID string, bookingReq Booking) error {
	booking := Booking{
		ID:       primitive.NewObjectID(),
		HostelID: bookingReq.HostelID,
		UserID:   userID,
		Date:     bookingReq.Date,
	}
	err := hostel.BookHostel(booking.HostelID)
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

func getBookings(userID string) ([]BookingDetail, error) {
	bookings := []Booking{}
	bookingsDetail := []BookingDetail{}
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

	for _, b := range bookings {
		hostel, _ := hostel.GetHostelByID(b.HostelID)
		bookingDetail := BookingDetail{
			BookingID:  b.ID.Hex(),
			HostelID:   hostel.ID.Hex(),
			Price:      hostel.Price,
			HostelName: hostel.Name,
			Date:       b.Date,
		}
		bookingsDetail = append(bookingsDetail, bookingDetail)
	}

	return bookingsDetail, nil
}

func getBookingDetail(bookingID string) (*BookingDetail, error) {
	booking := Booking{}
	bookingObjectID := transformToObjectID(bookingID)
	err := bookingCollection.FindOne(context.TODO(), bson.M{"_id": bookingObjectID}).Decode(&booking)
	if err != nil {
		log.Printf("Error while getting a booking information.")
		return nil, err
	}
	hostel, _ := hostel.GetHostelByID(booking.HostelID)
	bookingDetail := BookingDetail{
		BookingID:  booking.ID.Hex(),
		HostelID:   hostel.ID.Hex(),
		Price:      hostel.Price,
		HostelName: hostel.Name,
		Date:       booking.Date,
	}
	return &bookingDetail, nil
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

func transformToObjectID(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)
	return objectID
}
