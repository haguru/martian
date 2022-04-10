package mongo

import (
	"context"

	"github.com/haguru/martian/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ClientMongoInterface interface {
	Close(ctx context.Context) error
	TestConnection(ctx context.Context, readPerformance *readpref.ReadPref) error
	GetConnection() *mongo.Client
}
type MongoInterface interface {
	GetCollection(databaseName string, collectionName string) *mongo.Collection
	InsertDocument(ctx context.Context, collection *mongo.Collection, document models.Document)
	InsertDocuments(ctx context.Context, collection *mongo.Collection, documents []models.Document)
	UpdateDocument(ctx context.Context, collection *mongo.Collection, filter primitive.D, update primitive.D)
	FindDocuments(ctx context.Context, collection *mongo.Collection, filter primitive.D, options *options.FindOptions)
}
