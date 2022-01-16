package main

import (
	"context"
	"players/handler"
	pb "players/proto"
	"players/repository"
	"time"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	service = "players"
	version = "latest"
)

func main() {
	r := DBSetup()
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	p := handler.Players{
		MD: r,
	}
	// Register handler
	pb.RegisterPlayersHandler(srv.Server(), &p)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func DBSetup() repository.IRepository {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	database := client.Database("go-micro")
	playersCollection := database.Collection("playersDB")
	return repository.NewMongoDB(*playersCollection)
}
