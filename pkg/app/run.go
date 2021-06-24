package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
