package miniurls

import "github.com/DiegoSepuSoto/mini-url-service/src/shared"

type miniURLsRepository struct {
	mongoDBCollection shared.MongoDBCollection
}

func NewMongoDBMiniURLsRepository(mongoDBCollection shared.MongoDBCollection) *miniURLsRepository {
	return &miniURLsRepository{
		mongoDBCollection: mongoDBCollection,
	}
}
