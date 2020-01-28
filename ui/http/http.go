package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jschwinger23/aaec/config"
	"github.com/jschwinger23/aaec/ui"
)

func ListenAndServe() (err error) {
	conf := config.GetConfig()

	route := gin.Default()
	route.POST("/events", createEvent)
	return route.RunUnix(conf.UnixBind)
}

func createEvent(c *gin.Context) {
	req := Event{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	application, err := ui.GetApplication()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	event := struct {
		ID        string
		CreatedAt int64
		Package   string
		Type      string
	}{req.ID, req.CreatedAt, req.Package, req.Type}
	if err := application.CreateEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
