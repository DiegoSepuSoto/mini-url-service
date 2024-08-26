package miniurls

import (
	"context"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

func (r *miniURLsRepository) GetMinifiedURL(ctx context.Context, miniURL string) (string, error) {
	ctx, span := otel.Tracer(shared.TracerName).Start(ctx, "GetMinifiedURLRedisRepository")
	defer span.End()

	minifiedURL, err := r.redisClient.Get(ctx, miniURL).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			log.Printf("%s key does not exist on cache, moving to mongodb layer", miniURL)
		} else {
			log.Printf("error retrieving %s key from cache: %s, moving to mongodb layer", miniURL, err.Error())
		}

		return r.getMinifiedURLFromMongoDB(ctx, miniURL)
	}

	return minifiedURL, nil
}

func (r *miniURLsRepository) getMinifiedURLFromMongoDB(ctx context.Context, originalURL string) (string, error) {
	return r.mongoDBRepository.GetMinifiedURL(ctx, originalURL)
}
