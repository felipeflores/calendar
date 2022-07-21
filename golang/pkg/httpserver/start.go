package httpserver

import (
	"iot/ui"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	ui.AddRoutes(router)

	router.Run(":4000")
}
