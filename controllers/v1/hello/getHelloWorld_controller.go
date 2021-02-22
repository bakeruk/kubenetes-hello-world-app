package controllersv1hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// List out "hello world"
func List(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello world")
}
