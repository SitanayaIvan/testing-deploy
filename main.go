package main

import (
	"net/http"

	"github.com/SitanayaIvan/testing-deploy/config"
	"github.com/gin-gonic/gin"
)

func main() {

	conf, _ := config.LoadVariable()

	r := gin.Default()
	r.GET("api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!!",
		})
	})
	r.Run(":" + conf.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
