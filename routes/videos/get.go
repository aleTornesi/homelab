package videos

import (
	"database/sql"
	"fmt"
	"homelab/db"
	"homelab/frontend"
	"homelab/secrets"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	conn, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("HOMELAB_DB_HOST"),
			os.Getenv("HOMELAB_DB_PORT"),
			os.Getenv("HOMELAB_DB_USER"),
			secrets.GetDBPassword(),
			os.Getenv("HOMELAB_DB_NAME"),
			os.Getenv("HOMELAB_DB_SSLMODE"),
		),
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	parsedLimit, err := strconv.Atoi(limit)
	if err != nil && limit != "" {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	parsedOffset, err := strconv.Atoi(offset)
	if err != nil && offset != "" {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var videos []db.Video
	if limit != "" && offset != "" {
		videos, err = db.New(conn).GetVideosFromId(c, db.GetVideosFromIdParams{
			Limit: int32(parsedLimit),
			ID:    int32(parsedOffset),
		})
	} else {
		videos, err = db.New(conn).GetVideos(c, int32(parsedLimit))
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Type", "text/html")
	frontend.MediaCardList(videos, "primary").Render(c, c.Writer)
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	conn, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("HOMELAB_DB_HOST"),
			os.Getenv("HOMELAB_DB_PORT"),
			os.Getenv("HOMELAB_DB_USER"),
			secrets.GetDBPassword(),
			os.Getenv("HOMELAB_DB_NAME"),
			os.Getenv("HOMELAB_DB_SSLMODE"),
		),
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer conn.Close()

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	video, err := db.New(conn).GetVideoById(c, int32(parsedId))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Header("Content-Type", "text/html")
	frontend.VideoDetail(video).Render(c, c.Writer)
}
