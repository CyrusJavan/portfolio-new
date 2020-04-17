package blog

import (
	"net/http"

	"github.com/apex/log"

	"github.com/CyrusJavan/portfolio-new/src/db"
	"github.com/gin-gonic/gin"
)

type trackResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func handleAPITrack(c *gin.Context) {
	db := db.GetInstance(c)

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
		l := c.Keys["logEntry"].(*log.Entry)
		l.WithError(err).WithFields(log.Fields{
			"func": "handleAPITrack",
		}).Error("insert failed")
		c.JSON(http.StatusInternalServerError, trackResponse{"fail", err.Error()})
		return
	}

	c.JSON(http.StatusOK, trackResponse{"success", ""})
}
