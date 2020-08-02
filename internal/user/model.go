package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User stores user information
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Lastname string             `json:"lastname"`
	Password string             `json:"password"`
}

// UserWithoutPassword stores same information as user does but without password
type UserWithoutPassword struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	Lastname string             `json:"lastname"`
}

func (user User) removeUserPassword() *UserWithoutPassword {
	return &UserWithoutPassword{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Lastname: user.Lastname,
	}
}

// Booking stores booking information
type Booking struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	UserID   string             `json:"userId" bson:"userId"`
	HostelID string             `json:"hostelId" bson:"hostelId"`
	Date     string             `json:"date" bson:"date"`
}

// BookingDetail represents the booking information for user
type BookingDetail struct {
	BookingID  string  `json:"bookingId"`
	HostelID   string  `json:"hostelId"`
	HostelName string  `json:"hostelName"`
	Date       string  `json:"date"`
	Price      float32 `json:"price"`
}
