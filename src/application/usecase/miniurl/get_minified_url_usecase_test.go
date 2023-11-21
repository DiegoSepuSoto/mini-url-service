package miniurl

import (
	"context"
	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestGetMinifiedURLUseCase(t *testing.T) {
	t.Run("when get minified url use case is executed as expected retrieving existing minified url", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return("https://www.google.cl", nil)

		miniURLUseCase := NewMiniURLUseCase(mockMiniURLRepository)

		miniURLResponse, err := miniURLUseCase.GetMinifiedURL(context.Background(), "/abc123")

		mockMiniURLRepository.AssertNotCalled(t, "CreateNewMiniURL")

		assert.Nil(t, err)
		assert.Equal(t, "https://www.google.cl", miniURLResponse.MinifiedURL)
	})

	t.Run("when get minified url use case gets an unwanted error from repository", func(t *testing.T) {
		mockMiniURLRepository := new(miniURLsRepositoryMock)

		mockMiniURLRepository.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return("", shared.BuildError(http.StatusNotFound, shared.DatabaseNotFoundError, "not found in db", "miniURLsRepository"))

		miniURLUseCase := NewMiniURLUseCase(mockMiniURLRepository)

		miniURLResponse, err := miniURLUseCase.GetMinifiedURL(context.Background(), "/abc123")

		assert.NotNil(t, err)
		assert.Nil(t, miniURLResponse)
		assert.Equal(t, http.StatusNotFound, err.(*shared.ApplicationError).HTTPStatusCode)
		assert.Equal(t, shared.DatabaseNotFoundError, err.(*shared.ApplicationError).ErrorCode)
		assert.Equal(t, "not found in db", err.(*shared.ApplicationError).ErrorDescription)
		assert.Equal(t, "miniURLsRepository", err.(*shared.ApplicationError).ErrorOrigin)
	})
}
