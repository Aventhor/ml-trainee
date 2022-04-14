package api

import (
	"server/cmd/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app * gin.Engine) {
	app.POST("/", handlers.PoseDetect)
}