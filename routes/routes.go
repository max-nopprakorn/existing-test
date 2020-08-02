package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/existing-test/helper"
	authHandler "github.com/existing-test/internal/auth"
	hostelHanlder "github.com/existing-test/internal/hostel"
	userHandler "github.com/existing-test/internal/user"
)

// Routes will handle every routing for this application
func Routes(router *gin.Engine) {
	router.GET("/", emptyRoute)

	hostel := router.Group("/hostels")
	hostel.Use(helper.JWTAuthMiddleware())
	{
		hostel.GET("/", hostelHanlder.GetHostelsHandler)
		hostel.GET("/:hostelId", hostelHanlder.GetHostelByIDHandler)
		hostel.POST("/", hostelHanlder.CreateHostelHandler)
	}

	user := router.Group("/user")
	user.Use(helper.JWTAuthMiddleware())
	{
		user.GET("/", userHandler.GetUserDetailHandler)
		user.GET("/bookings", userHandler.GetBookingsHandler)
		user.GET("/bookings/:bookingId", userHandler.GetBookingDetailHandler)
		user.POST("/book", userHandler.BookHostelHandler)
	}

	router.POST("/register", authHandler.RegisterHandler)
	router.POST("/login", authHandler.LoginHandler)
	router.NoRoute(noRoute)

}
func noRoute(c *gin.Context) {
	c.AbortWithStatusJSON(404, gin.H{
		"message": "Path not found.",
	})
}

func emptyRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Application.",
	})
}
