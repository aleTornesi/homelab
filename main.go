package main

import (
	"homelab/frontend"

	"homelab/routes/photos"
	"homelab/routes/songs"
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
	r.GET("/media/videos/:id", videos.GetById)

	r.GET("/media/photos", photos.Get)
	r.GET("/media/photos/:id", photos.GetById)
	r.GET("/media/photos/:id/view", photos.Serve)

	r.GET("/media/songs", songs.Get)
	r.GET("/media/songs/:id", songs.GetById)

	r.Run(":8080")
}


