package main

import (
	"github.com/bakeruk/kubernetes-hello-world-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialises Gin
	router := gin.Default()
	// Initialises the routes
	routes.InitAllRoutes(router)
	// listen and serve on 0.0.0.0:3000 (for windows "localhost:3000")
	router.Run(":3000")
}
