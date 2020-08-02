package main

import (
	"github.com/existing-test/config"
	"github.com/existing-test/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func main() {

	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)
	appengine.Main()
	// log.Fatal(router.Run(":8090"))
}
