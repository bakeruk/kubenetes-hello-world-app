package main

import (
	"github.com/bakeruk/kubernetes-hello-world-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialises Gin
	router := gin.Default()
	// Initialises the routes
	routes.InitAllRoutes(router)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run(":8080")
}
