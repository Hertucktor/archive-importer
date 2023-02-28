package mongodb

import (
	"context"
	"github.com/Hertucktor/archive-importer/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func CreateClient(logger *zap.SugaredLogger) (*mongo.Client, context.Context, context.CancelFunc, error) {
	conf, err := config.GetConfig("config.yml", logger)
	if err != nil {
		logger.Errorf("Error: couldn't receive env vars")
		return nil, nil, nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + conf.DBUser + ":" + conf.DBPass + "@" + conf.DBPort + "/" + conf.DBName))
	if err != nil {
		logger.Error(err)
		return client, nil, nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.Error(err)
		return client, ctx, cancelFunc, err
	}

	return client, ctx, cancelFunc, err
}
