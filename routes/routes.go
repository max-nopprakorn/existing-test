package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// router = gin.Default()

	hostel := router.Group("/hostels")
	{
		hostel.GET("/", hostelHanlder.GetHostels)
		hostel.GET("/:hostelID", hostelHanlder.GetHostelById)
		hostel.POST("/", hostelHanlder.CreateHostel)
	}

	user := router.Group("/user")
	{
		user.GET("/:userId", userHandler.GetUserById)
		user.GET("/:userId/bookings", userHandler.GetBookings)
		user.GET("/:userId/bookings/:bookingId", userHandler.GetBookingInfo)
	}

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

}
