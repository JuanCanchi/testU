package controllers

import (
	"net/http"
	"testU/internal/models"
	"testU/internal/services"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	service services.TweetService
}

func NewTweetController(service services.TweetService) *TweetController {
	return &TweetController{service: service}
}

func (tc *TweetController) CreateTweet(c *gin.Context) {
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.service.CreateTweet(&tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tweet created", "tweet": tweet})
}
