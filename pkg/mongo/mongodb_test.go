package mongo

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoDB_GetCollection(t *testing.T) {
	type fields struct {
		client ClientMongoInterface
	}
	type args struct {
		databaseName   string
		collectionName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *mongo.Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mongodb := &MongoDB{
				client: tt.fields.client,
			}
			if got := mongodb.GetCollection(tt.args.databaseName, tt.args.collectionName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoDB.GetCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
