package db

import (
	"context"
	"time"

	"github.com/0xRokib/golang_auth/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {

	clientOpts := options.Client().ApplyURI(cfg.Mongo_URI)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := client.Ping(connectCtx, nil); err != nil {
		return nil, err
	}

	db := client.Database(cfg.Mongo_DB)

	return &Mongo{
		Client: client,
		Db:     db,
	}, nil

}
