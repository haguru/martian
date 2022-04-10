package mongo

import (
	"context"

	"github.com/haguru/martian/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client ClientMongoInterface
}

func NewMongoDB(client ClientMongoInterface) MongoInterface {

	return &MongoDB{}
}

func (mongodb *MongoDB) GetCollection(databaseName string, collectionName string) *mongo.Collection {
	client := mongodb.client.GetConnection()
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}

func (mongodb *MongoDB) InsertDocument(ctx context.Context, collection *mongo.Collection, document models.Document) {

}

func (mongodb *MongoDB) InsertDocuments(ctx context.Context, collection *mongo.Collection, documents []models.Document){

}

func (mongodb *MongoDB) UpdateDocument(ctx context.Context, collection *mongo.Collection, filter primitive.D, update primitive.D){

}

func (mongodb *MongoDB) FindDocuments(ctx context.Context, collection *mongo.Collection, filter primitive.D, options *options.FindOptions){
	
}
