package routes

import (
	"github.com/gin-gonic/gin"

	hostelHanlder "github.com/existing-test/internal/hostel"
)

func Routes(router *gin.Engine) {
	// router = gin.Default()

	hostel := router.Group("/hostels")
	{
		hostel.GET("/", hostelHanlder.GetHostelsHandler)
		hostel.GET("/:hostelID", hostelHanlder.GetHostelByIdHandler)
		hostel.POST("/", hostelHanlder.CreateHostelHandler)
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
