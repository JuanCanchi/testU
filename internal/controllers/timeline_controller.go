package controllers

import (
	"net/http"
	"testU/internal/services"

	"github.com/gin-gonic/gin"
)

type TimelineController struct {
	service services.TweetService
}

func NewTimelineController(service services.TweetService) *TimelineController {
	return &TimelineController{service: service}
}

func (tc *TimelineController) GetTimeline(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	tweets, err := tc.service.GetTimeline(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch timeline"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Timeline fetched", "tweets": tweets})
}
