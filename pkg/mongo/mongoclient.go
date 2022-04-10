package mongo

import (
	"context"
	"fmt"

	"github.com/haguru/martian/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(mongoConfig config.MongoConfiguration) (ClientMongoInterface, error) {

	mongoURI := fmt.Sprintf("mongodb://%v:%d", mongoConfig.Host, mongoConfig.Port)
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to create mongo client: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("unable to ping mongo: %v", err)
	}

	mongoClient := &MongoClient{
		client: client,
	}

	return mongoClient, nil

}

func (mclient *MongoClient) Close(ctx context.Context) error {
	return mclient.client.Disconnect(ctx)
}

func (mclient *MongoClient) TestConnection(ctx context.Context, rp *readpref.ReadPref)  error{
	err := mclient.client.Ping(ctx, rp)

	if err != nil {
		return fmt.Errorf("unable to ping mongo: %v", err)
	}
	return nil
}
func (mclient *MongoClient) GetConnection() *mongo.Client{
	return mclient.client
}

