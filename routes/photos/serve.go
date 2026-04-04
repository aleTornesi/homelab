package photos

import (
	"database/sql"
	"fmt"
	"homelab/db"
	"homelab/secrets"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Serve(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

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

	photo, err := db.New(conn).GetPhotoById(c, int32(id))
	if err == sql.ErrNoRows {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	println(photo.Path)
	c.File(photo.Path)
}
