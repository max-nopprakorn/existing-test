package helper

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CreateToken will generate token from user id
func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetUserIDFromToken will extract token and return user id
func GetUserIDFromToken(c *gin.Context) string {
	tokenString := c.Request.Header.Get("X-Auth-Token")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	userID := claims["userId"].(string)
	return userID
}

// JWTAuthMiddleware is the middleware to check if token is valid or not
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized request.",
		})
		return
	} else if checkToken(token) {
		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized request.",
		})
	}
}

func checkToken(tokenString string) bool {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})

	if err != nil {
		log.Printf("Error while extracting token, because of %v", err)
		return false
	}

	authorized := claims["authorized"].(bool)
	if !authorized {
		log.Printf("This token is not authorized.")
		return false
	}

	return true
}
