package main

import (
	"net/http"
	"sns_api/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Worldab"})
	})

	r.Run(":" + config.Config.ServerPort)
}
