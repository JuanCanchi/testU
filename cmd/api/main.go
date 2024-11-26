package main

import (
	"log"
	"testU/internal/repositories"
	"testU/internal/services"
	"testU/pkg/db"
)

import (
	"github.com/gin-gonic/gin"
	"testU/internal/controllers"
)

func main() {
	db.InitMongo()

	if db.GetMongoClient() == nil {
		log.Fatalf("Error: No se pudo conectar a MongoDB")
	}

	tweetRepo := repositories.NewTweetRepository()
	followRepo := repositories.NewFollowRepository()

	tweetService := services.NewTweetService(tweetRepo, followRepo)
	followService := services.NewFollowService(followRepo)

	tweetController := controllers.NewTweetController(tweetService)
	followController := controllers.NewFollowController(followService)
	timelineController := controllers.NewTimelineController(tweetService)

	r := gin.Default()

	r.POST("/tweets", tweetController.CreateTweet)
	r.POST("/follow", followController.FollowUser)
	r.GET("/timeline", timelineController.GetTimeline)

	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
