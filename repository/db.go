package repository

import (
	"context"
	"fmt"
	players "players/proto"

	"go-micro.dev/v4/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISpecification interface {
	IsSatisfied(in interface{}) bool
}

type AgeSpecification struct {
	Age int
}

func (as *AgeSpecification) IsSatisfied(in interface{}) bool {
	player := in.(*players.Player)
	return as.Age < int(player.GetAge())
}

type IRepository interface {
	Save(ctx context.Context, entity interface{}) (interface{}, error)
	GetById(ctx context.Context, id string) (interface{}, error)
	Update(ctx context.Context, entity interface{}) (interface{}, error)
	Delete(ctx context.Context, entity interface{}) error
	GetAll(ctx context.Context) (interface{}, error)
	GetFilter(ctx context.Context, fk, fv interface{}) (interface{}, error)
	GetSpec(ctx context.Context, spec ISpecification) (interface{}, error)
}

type MongoDB struct {
	collection mongo.Collection
}

func NewMongoDB(c mongo.Collection) IRepository {
	return &MongoDB{collection: c}
}

func (md *MongoDB) Save(ctx context.Context, entity interface{}) (interface{}, error) {
	player, ok := entity.(*players.Player)
	if !ok {
		return nil, fmt.Errorf("invalid request")
	}
	res, err := md.collection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: player.Name},
		{Key: "Age", Value: player.Age},
		{Key: "Position", Value: player.Position},
		{Key: "ClubName", Value: player.ClubName},
	})
	if err != nil {
		logger.Infof("InsertOne error: %v", err)
		return nil, err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (md *MongoDB) GetById(ctx context.Context, id string) (interface{}, error) {
	var player players.Player
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Errorf("Invalid ID")
		return nil, err
	}
	if err := md.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&player); err != nil {
		logger.Errorf("marshelling error: %v", err)
		return nil, err
	}
	logger.Infof("From DB: %v", &player)
	return &player, nil
}

func (md *MongoDB) Update(ctx context.Context, entity interface{}) (interface{}, error) {
	player, ok := entity.(*players.UpdateRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request")
	}
	objId, err := primitive.ObjectIDFromHex(player.Id)
	if err != nil {
		logger.Errorf("primitive.ObjectIDFromHex() error: %v", err)
		return nil, err
	}
	_, err = md.collection.UpdateOne(
		ctx,
		bson.M{"_id": objId},
		bson.D{
			{"$set", bson.D{
				{Key: "Name", Value: player.Name},
				{Key: "Age", Value: player.Age},
				{Key: "Position", Value: player.Position},
				{Key: "ClubName", Value: player.ClubName},
			}},
		},
	)
	if err != nil {
		logger.Errorf("UpdateOne error: %v", err)
		return nil, err
	}
	return md.GetById(ctx, player.Id)
}

func (md *MongoDB) Delete(ctx context.Context, entity interface{}) error {
	req := entity.(*players.DeleteRequest)
	_, err := md.collection.DeleteOne(ctx, bson.M{
		"_id": req.Id,
	})
	if err != nil {
		logger.Errorf("Mongo DB DeleteOne error: %v", err)
		return err
	}
	return nil
}

func (md *MongoDB) GetAll(ctx context.Context) (interface{}, error) {
	cursor, err := md.collection.Find(ctx, bson.D{})
	if err != nil {
		logger.Errorf("Mongo Find error: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	var mPlayers []*players.Player
	for cursor.Next(ctx) {
		p := players.Player{}
		err := cursor.Decode(&p)
		logger.Infof("p: %v", &p)
		if err != nil {
			logger.Errorf("cursor.Decode() error: %v", err)
			return nil, err
		}
		mPlayers = append(mPlayers, &p)
	}
	return mPlayers, nil
}

func (md *MongoDB) GetFilter(ctx context.Context, fk, fv interface{}) (interface{}, error) {
	filterKey := fk.(string)
	filterValue := fv.(string)
	filterCursor, err := md.collection.Find(ctx, bson.M{filterKey: filterValue})
	if err != nil {
		logger.Errorf("Find() error: %v", err)
		return nil, err
	}
	defer filterCursor.Close(ctx)
	var mPlayers []*players.Player
	for filterCursor.Next(ctx) {
		p := players.Player{}
		err := filterCursor.Decode(&p)
		if err != nil {
			logger.Errorf("cursor.Decode() error: %v", err)
			return nil, err
		}
		mPlayers = append(mPlayers, &p)
	}
	return mPlayers, nil
}

func (md *MongoDB) GetSpec(ctx context.Context, spec ISpecification) (interface{}, error) {
	allPlayers, err := md.GetAll(ctx)
	logger.Infof("GetAll: %v", allPlayers)
	if err != nil {
		logger.Errorf("GetAll() error: %v", err)
		return nil, err
	}
	all := allPlayers.([]*players.Player)
	var filteredPlayers []*players.Player
	for _, v := range all {
		if spec.IsSatisfied(v) {
			filteredPlayers = append(filteredPlayers, v)
		}
	}
	return filteredPlayers, nil
}
