package main

import (
	"fmt"
	"net/http"

	"server/cmd/api"
	"server/cmd/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewConfig()

	app := gin.Default()
	app.StaticFS("/uploads", http.Dir(config.UploadsPath))

	api.SetupRoutes(app)

	app.Run(":" + fmt.Sprint(config.Port))
}

