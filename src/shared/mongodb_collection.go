package shared

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
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
	ctx, span := otel.Tracer(TracerName).Start(ctx, "MongoDBFindOne")
	defer span.End()

	result := m.mongoCollection.FindOne(ctx, filter, opts...)

	return result
}

func CreateMongoDBCollection() *mongoDBCollection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(fmt.Sprintf("could not connect to database: [%s]", err.Error()))
	}

	return &mongoDBCollection{
		mongoCollection: client.Database("marketingDB").Collection("miniurls"),
	}
}
