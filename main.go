package main

import (
	"homelab/frontend"
	"homelab/routes/videos"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		frontend.Home().Render(c, c.Writer)
	})

	r.GET("/media/videos", videos.Get)
	r.GET("/media/:id", videos.GetById)

	r.Run(":8080")
}


