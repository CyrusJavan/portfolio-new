package blog

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/src/db"
)

type trackResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func handleAPITrack(c *gin.Context) {
	db := db.GetInstance()

	viewType := c.PostForm("type")
	page := c.PostForm("page")

	_, err := db.NamedExec(`
	INSERT INTO view_track (page, view_type, ip_address, user_agent)
	VALUES (:page, :view_type, :ip_address, :user_agent)`,
		map[string]interface{}{
			"page":       page,
			"view_type":  viewType,
			"ip_address": c.ClientIP(),
			"user_agent": c.GetHeader("User-Agent"),
		})

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, trackResponse{"fail", err.Error()})
		return
	}

	c.JSON(http.StatusOK, trackResponse{"success", ""})
}
