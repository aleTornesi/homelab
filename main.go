package main

import (
	"homelab/frontend"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		frontend.Home().Render(c, c.Writer)
	})

	r.Run(":8080")

}


