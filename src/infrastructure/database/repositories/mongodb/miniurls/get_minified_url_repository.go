package miniurls

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

func (r *miniURLsRepository) GetMinifiedURL(ctx context.Context, miniURL string) (string, error) {
	filter := bson.D{{Key: "new_url", Value: miniURL}}

	var miniURLRecord *entities.MiniURLRecord
	err := r.mongoDBCollection.FindOne(ctx, filter).Decode(&miniURLRecord)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			log.Printf("%s key does not exist on mongodb, sending error", miniURL)
			return "", shared.BuildError(
				http.StatusNotFound,
				shared.DatabaseNotFoundError,
				err.Error(),
				"miniURLsRepository")
		}

		log.Printf("error finding %s key on mongodb, sending error", miniURL)
		return "", shared.BuildError(http.StatusInternalServerError, shared.DatabaseFindError, err.Error(), "miniURLsRepository")
	}

	return miniURLRecord.OriginalURL, nil
}
