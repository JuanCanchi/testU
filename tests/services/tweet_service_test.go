package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testU/internal/models"
	"testU/internal/services"
	"testing"
	"time"
)

// Mock de TweetRepository
type MockTweetRepository struct {
	mock.Mock
}

func (m *MockTweetRepository) CreateTweet(tweet *models.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

func (m *MockTweetRepository) GetTweetsByUserIDs(userIDs []string) ([]*models.Tweet, error) {
	args := m.Called(userIDs)
	return args.Get(0).([]*models.Tweet), args.Error(1)
}

func TestCreateTweet(t *testing.T) {
	mockRepo := new(MockTweetRepository)
	mockRepoFollow := new(MockFollowRepository)

	service := services.NewTweetService(mockRepo, mockRepoFollow)

	tweet := &models.Tweet{
		UserID:    "user1",
		Content:   "Hello, world!",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateTweet", tweet).Return(nil)

	err := service.CreateTweet(tweet)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetTweetsByUser(t *testing.T) {
	mockRepoTweet := new(MockTweetRepository)
	mockRepoFollow := new(MockFollowRepository)

	mockRepoFollow.On("GetFollowedUserIDs", "user1").Return([]string{"user2", "user3"}, nil)

	service := services.NewTweetService(mockRepoTweet, mockRepoFollow)
	now := time.Now()
	userID := "user1"
	tweets := []*models.Tweet{
		{UserID: "user2", Content: "Tweet 1", CreatedAt: now},
		{UserID: "user3", Content: "Tweet 2", CreatedAt: now},
	}

	mockRepoTweet.On("GetTweetsByUserIDs", []string{"user2", "user3"}).Return(tweets, nil)

	result, err := service.GetTimeline(userID)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Tweet 1", result[0].Content)
	assert.Equal(t, "Tweet 2", result[1].Content)

	mockRepoTweet.AssertExpectations(t)
	mockRepoFollow.AssertExpectations(t)
}
