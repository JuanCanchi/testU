package interfaces

import "testU/internal/models"

type TweetRepository interface {
	CreateTweet(tweet *models.Tweet) error
	GetTweetsByUserIDs(userIDs []string) ([]*models.Tweet, error)
}

type FollowRepository interface {
	SaveFollow(followerID, followedID string) error
	GetFollowedUserIDs(userID string) ([]string, error)
}
