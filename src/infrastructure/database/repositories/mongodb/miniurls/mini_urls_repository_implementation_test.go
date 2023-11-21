package miniurls

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/database/repositories/mongodb/miniurls/entities"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

func TestGetMinifiedURLRepository(t *testing.T) {
	t.Run("when get minified url gets executed as expected and finds the record in the database", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*entities.MiniURLRecord)

			arg.OriginalURL = "https://www.google.cl"
		}).Return(nil)

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMongoDBMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetMinifiedURL(context.Background(), "www.google.cl")

		assert.Nil(t, err)
		assert.Equal(t, "https://www.google.cl", miniURL)
	})

	t.Run("when get minified url gets executed as expected and does not find record in the database", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Return(errors.New("mongo: no documents in result"))

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMongoDBMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetMinifiedURL(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Equal(t, "", miniURL)
		assert.Equal(t, http.StatusNotFound, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseNotFoundError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "mongo: no documents in result", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})

	t.Run("when get minified url an error finding record then returns the error", func(t *testing.T) {
		mockMongoCollection := new(mongoCollectionMock)
		mockMongoSingleResult := new(mongoSingleResultMock)

		mockMongoSingleResult.On("Decode", mock.AnythingOfType("*entities.MiniURLRecord")).Return(errors.New("mongo: error"))

		mockMongoCollection.On("FindOne", mock.Anything, mock.Anything, mock.AnythingOfType("[]*options.FindOneOptions")).
			Return(mockMongoSingleResult)

		miniURLsRepository := NewMongoDBMiniURLsRepository(mockMongoCollection)

		miniURL, err := miniURLsRepository.GetMinifiedURL(context.Background(), "www.google.cl")

		assert.NotNil(t, err)
		assert.Equal(t, "", miniURL)
		assert.Equal(t, http.StatusInternalServerError, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseFindError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "mongo: error", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
