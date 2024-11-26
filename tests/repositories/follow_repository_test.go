package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFollowRepository struct {
	mock.Mock
}

func (m *MockFollowRepository) SaveFollow(followerID, followedID string) error {
	args := m.Called(followerID, followedID)
	return args.Error(0)
}

func (m *MockFollowRepository) GetFollowedUserIDs(userID string) ([]string, error) {
	args := m.Called(userID)
	return args.Get(0).([]string), args.Error(1)
}

func TestSaveFollow(t *testing.T) {
	mockRepo := new(MockFollowRepository)

	followerID := "user1"
	followedID := "user2"

	mockRepo.On("SaveFollow", followerID, followedID).Return(nil)

	err := mockRepo.SaveFollow(followerID, followedID)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetFollowedUserIDs(t *testing.T) {
	mockRepo := new(MockFollowRepository)

	userID := "user1"
	followedIDs := []string{"user2", "user3"}

	mockRepo.On("GetFollowedUserIDs", userID).Return(followedIDs, nil)

	result, err := mockRepo.GetFollowedUserIDs(userID)

	assert.NoError(t, err)

	assert.ElementsMatch(t, followedIDs, result)

	mockRepo.AssertExpectations(t)
}
