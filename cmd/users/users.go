package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("Error al crear el cliente de MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	usersCollection := client.Database("twitter_clone").Collection("users")

	userCount := 10
	var wg sync.WaitGroup
	wg.Add(userCount)

	orderChannel := make(chan struct{}, 1)

	for i := 1; i <= userCount; i++ {
		go func(userID int) {
			defer wg.Done()

			orderChannel <- struct{}{}

			user := bson.M{
				"_id":  fmt.Sprintf("user%d", userID),
				"name": fmt.Sprintf("User%d", userID),
			}

			_, err := usersCollection.InsertOne(ctx, user)
			if err != nil {
				log.Printf("Error al insertar el usuario %d: %v", userID, err)
			} else {
				log.Printf("Usuario %d creado exitosamente", userID)
			}

			<-orderChannel
		}(i)
	}

	wg.Wait()
	log.Println("Todos los usuarios fueron creados")
}
