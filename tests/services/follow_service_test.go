package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testU/internal/services"
	"testing"
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

type FollowServiceMock struct {
	mock.Mock
}

func (fs *FollowServiceMock) FollowUser(followerID, followedID string) error {
	args := fs.Called(followerID, followedID)
	return args.Error(0)
}

func TestFollowUser(t *testing.T) {
	mockRepo := new(MockFollowRepository)
	mockRepo.On("SaveFollow", "user1", "user2").Return(nil)

	service := services.NewFollowService(mockRepo)
	err := service.FollowUser("user1", "user2")

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
