package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testU/internal/interfaces"
	"testU/internal/models"
	"testU/pkg/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type tweetRepository struct{}

func NewTweetRepository() interfaces.TweetRepository {
	return &tweetRepository{}
}

func (tr *tweetRepository) CreateTweet(tweet *models.Tweet) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tweet.CreatedAt = time.Now()
	err := tr.ensureIndexes(ctx)
	if err != nil {
		log.Printf("Error al crear los índices: %v", err)
		return err
	}

	result, err := db.TweetsCollection.InsertOne(ctx, tweet)
	if err != nil {
		log.Printf("Error al insertar el tweet: %v", err)
		return err
	}
	tweet.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (tr *tweetRepository) GetTweetsByUserIDs(userIDs []string) ([]*models.Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": bson.M{"$in": userIDs}}
	cursor, err := db.TweetsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tweets []*models.Tweet
	if err := cursor.All(ctx, &tweets); err != nil {
		return nil, err
	}

	return tweets, nil
}

func (tr *tweetRepository) ensureIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "user_id", Value: 1},
			{Key: "created_at", Value: -1},
		},
		Options: options.Index().SetUnique(false),
	}

	_, err := db.TweetsCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Error al crear el índice: %v", err)
		return err
	}

	log.Println("Índice creado o ya existe.")
	return nil
}
