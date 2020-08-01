package routes

import (
	"github.com/gin-gonic/gin"

	hostelHanlder "github.com/existing-test/internal/hostel"
	userHandler "github.com/existing-test/internal/user"
)

func Routes(router *gin.Engine) {
	// router = gin.Default()

	hostel := router.Group("/hostels")
	{
		hostel.GET("/", hostelHanlder.GetHostelsHandler)
		hostel.GET("/:hostelID", hostelHanlder.GetHostelByIDHandler)
		hostel.POST("/", hostelHanlder.CreateHostelHandler)
	}

	user := router.Group("/user")
	{
		user.GET("/:userId", userHandler.GetUserDetailHandler)
		user.GET("/:userId/bookings", userHandler.GetBookingsHandler)
		user.GET("/:userId/bookings/:bookingId", userHandler.GetBookingDetailHandler)
	}

	router.POST("/book", userHandler.BookHostelHandler)

	// router.POST("/register", authHandler.Register)
	// router.POST("/login", authHandler.Login)

}
