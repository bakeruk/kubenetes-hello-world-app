package controllersv1hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// List out "hello world"
func List(context *gin.Context) {
	data := map[string]string{"message": "Hello world"}
	context.JSON(http.StatusOK, data)
}
