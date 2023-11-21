package miniurls

import (
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

type miniURLsRepository struct {
	mongoDBRepository repositories.MiniURLsRepository
	redisClient       shared.RedisClient
}

func NewRedisMiniURLsRepository(mongoDBRepository repositories.MiniURLsRepository, redisClient shared.RedisClient) *miniURLsRepository {
	return &miniURLsRepository{
		mongoDBRepository: mongoDBRepository,
		redisClient:       redisClient,
	}
}
