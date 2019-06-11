package http

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/hieuphq/califit/src/interfaces/http/v1"
	"github.com/hieuphq/califit/src/registry"
)

// Apply route
func Apply(app *gin.Engine, ctn registry.Container) {
	root := app.Group("/")

	v1.Apply(root, ctn)
}
