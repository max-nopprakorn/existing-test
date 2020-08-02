package auth

import (
	userHelper "github.com/existing-test/internal/user"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	loginFrom := LoginRequest{}
	err := c.ShouldBindJSON(&loginFrom)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Invalid json request.",
		})
		return
	}
	isPasswordCorrect := comparePassword(loginFrom)
	if !isPasswordCorrect {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Username or/and Password are not correct.",
		})
		return
	}
	token, err := login(loginFrom)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when generating token.",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func RegisterHandler(c *gin.Context) {
	user := userHelper.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "Invalid json request.",
		})
	}
	isDuplicate := checkDuplicateUser(user.Username)
	if isDuplicate {
		c.AbortWithStatusJSON(409, gin.H{
			"message": "Username already exists.",
		})
		return
	}
	err = register(user)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Something went wrong when creating a new user.",
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Register successfully.",
	})
}
