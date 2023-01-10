package db

import (
	"context"
	"fmt"

	"github.com/Graphql/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectToDB(cfg *config.Config) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MONGOHost, cfg.MONGOPort)))
}