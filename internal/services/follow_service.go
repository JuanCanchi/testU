package services

import "testU/internal/interfaces"

type FollowService interface {
	FollowUser(followerID, followedID string) error
}

type followService struct {
	repo interfaces.FollowRepository
}

func NewFollowService(repo interfaces.FollowRepository) FollowService {
	return &followService{repo: repo}
}

func (fs *followService) FollowUser(followerID, followedID string) error {
	return fs.repo.SaveFollow(followerID, followedID)
}
