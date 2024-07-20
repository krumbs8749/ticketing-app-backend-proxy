package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func GetEvents(c *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/api/events/getAll") // Update with the actual Spring backend URL

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}

func GetEventByID(c *gin.Context) {
	eventID := c.Param("id")
	client := resty.New()

	resp, err := client.R().
		Get("http://localhost:8080/api/events/get/" + eventID) // Update with the actual Spring backend URL

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode(), resp.Header().Get("Content-Type"), resp.Body())
}
