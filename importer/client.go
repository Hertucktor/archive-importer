package importer

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ImporterClient(dbUser, dbPass, dbPort, dbName string) (*mongo.Client, context.Context, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPass + "@" + dbPort + "/" + dbName))
	if err != nil {
		return client, nil, err
	}

	ctx := context.Background()
	if err = client.Connect(ctx); err != nil {
		return client, ctx, err
	}

	return client, ctx, err
}
