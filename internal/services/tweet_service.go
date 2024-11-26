package services

import (
	"testU/internal/interfaces"
	"testU/internal/models"
)

type TweetService interface {
	CreateTweet(tweet *models.Tweet) error
	GetTimeline(userID string) ([]*models.Tweet, error)
}

type tweetService struct {
	tweetRepo  interfaces.TweetRepository
	followRepo interfaces.FollowRepository
}

func NewTweetService(tweetRepo interfaces.TweetRepository, followRepo interfaces.FollowRepository) TweetService {
	return &tweetService{tweetRepo: tweetRepo, followRepo: followRepo}
}

func (ts *tweetService) CreateTweet(tweet *models.Tweet) error {
	return ts.tweetRepo.CreateTweet(tweet)
}

func (ts *tweetService) GetTimeline(userID string) ([]*models.Tweet, error) {
	followedUserIDs, err := ts.followRepo.GetFollowedUserIDs(userID)
	if err != nil {
		return nil, err
	}

	return ts.tweetRepo.GetTweetsByUserIDs(followedUserIDs)
}
