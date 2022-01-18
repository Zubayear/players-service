package handler

import (
	"context"
	"go-micro.dev/v4/logger"
	players "players/proto"
	"players/repository"
)

type Players struct {
	MD repository.IRepository
}

func (p *Players) Save(ctx context.Context, req *players.SaveRequest, rsp *players.SaveResponse) error {
	logger.Infof("Received Players.Save request %v", req)
	player := &players.Player{
		Name:     req.Name,
		Position: req.Position,
		ClubName: req.ClubName,
		Age:      req.Age,
	}
	val, err := p.MD.Save(ctx, player)
	if err != nil {
		logger.Infof("DB error: %v", err)
		return err
	}
	logger.Infof("saved value: %v", val)
	rsp.Id = val.(string)
	return nil
}

func (p *Players) Get(ctx context.Context, req *players.PlayerRequest, rsp *players.PlayerResponse) error {
	logger.Infof("Request received Players.Get %v", req)
	val, err := p.MD.GetById(ctx, req.Id)
	if err != nil {
		logger.Infof("DB error: %v", err)
		return err
	}
	rsp.Player = val.(*players.Player)
	return nil
}

func (p *Players) Update(ctx context.Context, req *players.UpdateRequest, rsp *players.UpdateResponse) error {
	logger.Infof("Request received Players.Update %v", req)
	val, err := p.MD.Update(ctx, req)
	if err != nil {
		logger.Infof("DB error: %v", err)
		return err
	}
	logger.Infof("saved value: %v", val)
	rsp.Player = val.(*players.Player)
	return nil
}

func (p *Players) GetAll(ctx context.Context, req *players.PlayersRequest, rsp *players.PlayersResponse) error {
	logger.Infof("Received Players.GetAll request %v", req)
	res, err := p.MD.GetAll(ctx)
	if err != nil {
		logger.Errorf("GetAll error: %v", err)
		return err
	}
	logger.Infof("Response: %v", res)
	rsp.Players = res.([]*players.Player)
	return nil
}

func (p *Players) Delete(ctx context.Context, req *players.DeleteRequest, rsp *players.DeleteResponse) error {
	logger.Infof("Request received Players.Delete %v", req)
	err := p.MD.Delete(ctx, req)
	if err != nil {
		logger.Errorf("Delete error: %v", err)
		return err
	}
	return nil
}

func (p *Players) Filter(ctx context.Context, req *players.FilterRequest, rsp *players.PlayersResponse) error {
	logger.Infof("Request received Players.Filter %v", req)
	res, err := p.MD.GetSpec(ctx, &repository.ClubNameSpecification{
		ClubName: req.FilterKey,
	})
	if err != nil {
		logger.Errorf("Filter error: %v", err)
		return err
	}
	logger.Infof("Response: %v", res)
	rsp.Players = res.([]*players.Player)
	return nil
}
