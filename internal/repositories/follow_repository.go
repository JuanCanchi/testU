package repositories

import (
	"context"
	"testU/internal/interfaces"
	"testU/pkg/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type followRepository struct{}

func NewFollowRepository() interfaces.FollowRepository {
	return &followRepository{}
}

func (r *followRepository) SaveFollow(followerID, followedID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	follow := bson.M{"follower_id": followerID, "followed_id": followedID}
	_, err := db.FollowsCollection.InsertOne(ctx, follow)
	return err
}

func (r *followRepository) GetFollowedUserIDs(userID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"follower_id": userID}
	cursor, err := db.FollowsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var follows []bson.M
	if err := cursor.All(ctx, &follows); err != nil {
		return nil, err
	}

	var userIDs []string
	for _, follow := range follows {
		followedID, ok := follow["followed_id"].(string)
		if ok {
			userIDs = append(userIDs, followedID)
		}
	}

	return userIDs, nil
}
