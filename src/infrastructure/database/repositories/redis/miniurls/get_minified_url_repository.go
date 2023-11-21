package miniurls

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/redis/go-redis/v9"
)

func (r *miniURLsRepository) GetMinifiedURL(ctx context.Context, miniURL string) (string, error) {
	minifiedURL, err := r.redisClient.Get(ctx, miniURL).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("%s key does not exist on cache, moving to mongodb layer", miniURL)
		} else {
			log.Printf("error retrieving %s key from cache: %s", miniURL, err.Error())
		}

		return r.getMinifiedURLFromMongoDB(ctx, miniURL)
	}

	return minifiedURL, nil
}

func (r *miniURLsRepository) getMinifiedURLFromMongoDB(ctx context.Context, originalURL string) (string, error) {
	return r.mongoDBRepository.GetMinifiedURL(ctx, originalURL)
}
