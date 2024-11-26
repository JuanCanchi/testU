package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	UsersCollection   *mongo.Collection
	TweetsCollection  *mongo.Collection
	FollowsCollection *mongo.Collection

	MongoClient *mongo.Client

	once sync.Once
)

func InitMongo() {
	once.Do(func() {
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
		if err != nil {
			log.Fatalf("Failed to create MongoDB client: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Connect(ctx); err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		MongoClient = client
		UsersCollection = client.Database("twitter_clone").Collection("users")
		TweetsCollection = client.Database("twitter_clone").Collection("tweets")
		FollowsCollection = client.Database("twitter_clone").Collection("follows")
		log.Println("Connected to MongoDB")
	})

}

func GetMongoClient() *mongo.Client {
	return MongoClient
}
