package user

import (
	"github.com/existing-test/internal/hostel"
	"github.com/gin-gonic/gin"
)

// GetUserDetailHandler will return the user information
func GetUserDetailHandler(c *gin.Context) {
	userID := c.Param("userId")
	user, err := getUserDetail(userID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when trying to query a user information,",
		})
		return
	}
	if user == (&User{}) {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "Resource not found.",
		})
		return
	}

	c.JSON(200, user)
}

// BookHostelHandler will handle when user book a hostel
func BookHostelHandler(c *gin.Context) {
	userID := c.Param("userId")
	hostelID := c.Param("hostelId")
	isAvailable := hostel.CheckIfAvaliable(hostelID)
	if !isAvailable {
		c.AbortWithStatusJSON(409, gin.H{
			"message": "This hostel is not available.",
		})
		return
	}
	err := bookHostel(userID, hostelID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when trying to book a hostel.",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Book a hostel successfully.",
	})
}

// GetBookingsHandler will return user's bookings
func GetBookingsHandler(c *gin.Context) {
	userID := c.Param("userId")
	bookings, err := getBookings(userID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when trying to query user's bookings.",
		})
		return
	}
	c.JSON(200, bookings)
}

// GetBookingDetailHandler will return the information of booking
func GetBookingDetailHandler(c *gin.Context) {
	bookingID := c.Param("bookingId")
	booking, err := getBookingDetail(bookingID)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when trying to query a booking detail.",
		})
		return
	}
	c.JSON(200, booking)
}
