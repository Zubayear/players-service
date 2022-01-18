package main

import (
	"context"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
	"players/handler"
	pb "players/proto"
	"players/repository"
	"strconv"
	"time"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
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
	host, err := loadConfig()

	if err != nil {
		log.Errorf("Config loading error: %v", err)
	}

	str := "mongodb://" + host.Address + ":" + strconv.Itoa(host.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(str))
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

	database := client.Database(host.Name)
	playersCollection := database.Collection(host.Collection)
	return repository.NewMongoDB(*playersCollection)
}

type Host struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
	Collection string `json:"collection"`
}

func loadConfig() (*Host, error) {
	enc := yaml.NewEncoder()
	c, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	if err != nil {
		log.Errorf("Error loading config %v", err)
		return nil, err
	}

	err = c.Load(file.NewSource(file.WithPath("./config.yaml")))
	if err != nil {
		return nil, err
	}

	host := &Host{}
	if err := c.Get("hosts", "database").Scan(host); err != nil {
		return nil, err
	}

	return host, nil
}
