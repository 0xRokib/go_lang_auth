package app

import (
	"context"
	"time"

	"github.com/0xRokib/golang_auth/internal/config"
	"github.com/0xRokib/golang_auth/internal/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.LoadEnv()
	if err != nil {
		return nil, err
	}

	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.Db,
	}, nil

}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return err
	}

	return nil
}
