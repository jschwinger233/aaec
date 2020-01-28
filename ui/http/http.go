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
	route.POST("/events", createEvents)
	return route.RunUnix(conf.UnixBind)
}

func createEvents(c *gin.Context) {
	req := Events{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	application, err := ui.GetApplication()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	events := []struct {
		ID        string
		CreatedAt int64
		Type      string
		Content   []byte
	}{}
	for _, e := range req.Events {
		events = append(events, struct {
			ID        string
			CreatedAt int64
			Type      string
			Content   []byte
		}{e.ID, e.CreatedAt, e.Type, e.Content})
	}
	if err := application.CreateEvents(events); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
