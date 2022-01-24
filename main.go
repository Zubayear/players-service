package main

import (
	"context"
	"players/handler"
	pb "players/proto"
	"players/repository"
	"strconv"
	"time"

	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
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

	// create a publisher
	pub1 := micro.NewPublisher("club.topic.pubsub.1", srv.Client())

	// Register handler
	pb.RegisterPlayersHandler(srv.Server(), &p)

	go sendEv("club.topic.pubsub.1", r, pub1)

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

// Send using the publisher
func sendEv(topic string, r repository.IRepository, p micro.Publisher) {
	t := time.NewTicker(10 * time.Second)

	for _ = range t.C {
		ev, err := r.GetById(context.Background(), "61e6e44d7feb9029a24ea875")

		if err != nil {
			log.Errorf(err.Error())
		}

		log.Infof("publishing %+v", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Infof("error publishing: %v", err)
		}
	}
}
