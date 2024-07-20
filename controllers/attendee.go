package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func RegisterAttendee(c *gin.Context) {
	eventID := c.Param("eventId")
	var attendee map[string]interface{}

	if err := c.ShouldBindJSON(&attendee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()

	resp, err := client.R().
		SetBody(attendee).
		Put("http://localhost:8080/api/events/update/attendees/" + eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
