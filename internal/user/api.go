package user

import (
	"github.com/existing-test/internal/hostel"
	"github.com/gin-gonic/gin"
)

func GetUserDetailHandler(c *gin.Context) {
	userID := c.Param("userId")
	user, err := getUserDetail(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when trying to query a user information,",
		})
		return
	}
	if user == (&User{}) {
		c.JSON(404, gin.H{
			"message": "Resource not found.",
		})
		return
	}

	c.JSON(200, user)
}

func BookHostelHandler(c *gin.Context) {
	userID := c.Param("userId")
	hostelID := c.Param("hostelId")
	isAvailable := hostel.CheckIfAvaliable(hostelID)
	if !isAvailable {
		c.JSON(409, gin.H{
			"message": "This hostel is not available.",
		})
		return
	}
	err := bookHostel(userID, hostelID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when trying to book a hostel.",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Book a hostel successfully.",
	})
}

func GetBookingsHandler(c *gin.Context) {
	userID := c.Param("userId")
	bookings, err := getBookings(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when trying to query user's bookings.",
		})
		return
	}
	c.JSON(200, bookings)
}

func GetBookingDetailHandler(c *gin.Context) {
	bookingID := c.Param("bookingId")
	booking, err := getBookingDetail(bookingID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong when trying to query a booking detail.",
		})
		return
	}
	c.JSON(200, booking)
}
