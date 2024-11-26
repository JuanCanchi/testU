package repositories

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testU/internal/models"
)

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

	tweet := &models.Tweet{
		UserID:    "user1",
		Content:   "Hello, world!",
		CreatedAt: time.Now(),
	}

	mockRepo.On("CreateTweet", tweet).Return(nil)

	err := mockRepo.CreateTweet(tweet)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetTweetsByUserIDs(t *testing.T) {
	mockRepo := new(MockTweetRepository)
	now := time.Now()
	userIDs := []string{"user1", "user2"}
	tweets := []*models.Tweet{
		{UserID: "user1", Content: "Tweet 1", CreatedAt: now},
		{UserID: "user2", Content: "Tweet 2", CreatedAt: now},
	}

	mockRepo.On("GetTweetsByUserIDs", userIDs).Return(tweets, nil)

	result, err := mockRepo.GetTweetsByUserIDs(userIDs)

	assert.NoError(t, err)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Tweet 1", result[0].Content)
	assert.Equal(t, "Tweet 2", result[1].Content)

	mockRepo.AssertExpectations(t)
}
