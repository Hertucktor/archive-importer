package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func CreateClient(dbUser, dbPass, dbPort, dbName string, logger *zap.SugaredLogger) (*mongo.Client, error) {
	var ctx = context.TODO()

	logger.Infof("%v%v%v%v", dbUser, dbPass, dbPort, dbName)
	clientOptions := options.Client().ApplyURI("mongodb://" + dbUser + ":" + dbPass + "@" + dbPort + "/" + dbName)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return client, err
}

func pingDB(client *mongo.Client, logger *zap.SugaredLogger) {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal(err)
	}
}

func checkDatabase(dbName string, client *mongo.Client, logger *zap.SugaredLogger) {
	logger.Info(client.Database(dbName, nil))
}
