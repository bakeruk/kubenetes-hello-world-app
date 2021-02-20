package routesv1

import (
	"github.com/gin-gonic/gin"
)

// InitV1Routes initialises /v1 routes
func InitV1Routes(router *gin.Engine) {
	// Creates the /v1 routerGroup
	v1RouteGroup := router.Group("/v1")

	// Initialise /v1/hello routes
	initV1HelloRoutes(v1RouteGroup)
}
