package main

import (
	"log"

	"github.com/existing-test/config"
	"github.com/existing-test/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8090"))
}
