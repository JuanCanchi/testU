package controllers

import (
	"net/http"
	"testU/internal/services"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	service services.FollowService
}

func NewFollowController(service services.FollowService) *FollowController {
	return &FollowController{service: service}
}

func (fc *FollowController) FollowUser(c *gin.Context) {
	var follow struct {
		UserID   string `json:"user_id" binding:"required"`
		FollowID string `json:"follow_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&follow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := fc.service.FollowUser(follow.UserID, follow.FollowID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User followed"})
}
