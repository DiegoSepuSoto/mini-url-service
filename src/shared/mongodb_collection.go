package shared

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCollection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResult
}

type mongoDBCollection struct {
	mongoCollection *mongo.Collection
}

type MongoDBSingleResult interface {
	Decode(v interface{}) error
}

func (m *mongoDBCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) MongoDBSingleResult {
	return m.mongoCollection.FindOne(ctx, filter, opts...)
}

func CreateMongoDBCollection() *mongoDBCollection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(fmt.Sprintf("could not connect to database: [%s]", err.Error()))
	}

	log.Println("mongodb connection successful!")

	return &mongoDBCollection{
		mongoCollection: client.Database("marketingDB").Collection("mini-urls"),
	}
}
