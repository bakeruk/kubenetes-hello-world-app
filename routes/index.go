package routes

import (
	routesV1 "github.com/bakeruk/kubernetes-hello-world-app/routes/v1"
	"github.com/gin-gonic/gin"
)

// InitRoutes initialises all routes
func InitRoutes(router *gin.Engine) {
	// Initialises /v1 routes
	routesV1.InitV1Routes(router)
}
