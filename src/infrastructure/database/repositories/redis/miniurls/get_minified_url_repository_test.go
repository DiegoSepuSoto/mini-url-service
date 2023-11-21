package miniurls

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMinifiedURLRepository(t *testing.T) {
	t.Run("when get minified url repository gets executed as expected sending information from cache layer", func(t *testing.T) {
		mockRedisStringResult := new(redisStringResultMock)
		mockRedisStringResult.On("Result").Return("https://www.google.com", nil)

		mockRedisClient := new(redisClientMock)
		mockRedisClient.On("Get", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRedisStringResult)

		mockMongoRepository := new(mongoDBRepositoryMock)

		redisMiniURLsRepository := NewRedisMiniURLsRepository(mockMongoRepository, mockRedisClient)

		minifiedURL, err := redisMiniURLsRepository.GetMinifiedURL(context.Background(), "abc123")

		assert.Nil(t, err)
		assert.Equal(t, "https://www.google.com", minifiedURL)
	})

	t.Run("when get minified url repository gets executed as expected sending information from database layer", func(t *testing.T) {
		mockRedisStringResult := new(redisStringResultMock)
		mockRedisStringResult.On("Result").Return("", errors.New("redis: nil"))

		mockRedisClient := new(redisClientMock)
		mockRedisClient.On("Get", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRedisStringResult)

		mockMongoRepository := new(mongoDBRepositoryMock)
		mockMongoRepository.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return("https://www.google.com", nil)

		redisMiniURLsRepository := NewRedisMiniURLsRepository(mockMongoRepository, mockRedisClient)

		minifiedURL, err := redisMiniURLsRepository.GetMinifiedURL(context.Background(), "abc123")

		assert.Nil(t, err)
		assert.Equal(t, "https://www.google.com", minifiedURL)
	})

	t.Run("when get minified url repository gets an error from cache layer", func(t *testing.T) {
		mockRedisStringResult := new(redisStringResultMock)
		mockRedisStringResult.On("Result").Return("", errors.New("redis error"))

		mockRedisClient := new(redisClientMock)
		mockRedisClient.On("Get", mock.Anything, mock.AnythingOfType("string")).
			Return(mockRedisStringResult)

		mockMongoRepository := new(mongoDBRepositoryMock)
		mockMongoRepository.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return("https://www.google.com", nil)

		redisMiniURLsRepository := NewRedisMiniURLsRepository(mockMongoRepository, mockRedisClient)

		minifiedURL, err := redisMiniURLsRepository.GetMinifiedURL(context.Background(), "abc123")

		assert.Nil(t, err)
		assert.Equal(t, "https://www.google.com", minifiedURL)
	})
}
