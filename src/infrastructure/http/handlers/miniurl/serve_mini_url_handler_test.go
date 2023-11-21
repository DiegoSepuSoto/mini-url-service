package miniurl

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/DiegoSepuSoto/mini-url-service/src/domain/models"
)

func TestServeMinifiedURLHandler(t *testing.T) {
	t.Run("when serve minified url handler is executed and returns as expected", func(t *testing.T) {
		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/abc123", nil)

		c := e.NewContext(req, rec)

		mockMiniURLUseCase := new(miniURLUseCaseMock)

		mockMiniURLUseCase.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return(&models.MinifiedURLResponse{MinifiedURL: "www.google.cl"}, nil)

		miniURLHandler := NewMiniURLHandler(e, mockMiniURLUseCase)

		err := miniURLHandler.ServeMiniURLHandler(c)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusMovedPermanently, rec.Code)
	})

	t.Run("when serve minified url handler sends an error from use case", func(t *testing.T) {
		e := echo.New()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/api/abc123", nil)
		req.Header.Set("Authorization", "Bearer xyz789")

		c := e.NewContext(req, rec)

		mockMiniURLUseCase := new(miniURLUseCaseMock)

		mockMiniURLUseCase.On("GetMinifiedURL", mock.Anything, mock.AnythingOfType("string")).
			Return(nil, errors.New("internal logic error"))

		miniURLHandler := NewMiniURLHandler(e, mockMiniURLUseCase)

		err := miniURLHandler.ServeMiniURLHandler(c)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusTemporaryRedirect, rec.Code)
	})
}
