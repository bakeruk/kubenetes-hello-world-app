package routesv1

import (
	helloController "github.com/bakeruk/kubernetes-hello-world-app/controllers/v1/hello"
	"github.com/gin-gonic/gin"
)

// Initialises /v1/hello routes
func initV1HelloRoutes(v1RouteGroup *gin.RouterGroup) {
	// Creates the /v1/hello routerGroup
	helloRouterGroup := v1RouteGroup.Group("/hello")

	/**
	 * @api {get} /v1/hello List
	 */
	helloRouterGroup.GET("/", helloController.List)
}
