package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func testHandlerWithURLParam(c echo.Context) error {
	return c.String(http.StatusOK, "response ok")
}

func TestURLParamMiddleware(t *testing.T) {
	t.Run("when url param middleware accepts a request with correct param", func(t *testing.T) {
		e := echo.New()
		e.Use(URLParamMiddleware)
		e.GET("/:mini-url", testHandlerWithURLParam)
		req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "response ok", rec.Body.String())
	})

	t.Run("when url param middleware rejects a request with no param", func(t *testing.T) {
		e := echo.New()
		e.Use(URLParamMiddleware)
		e.GET("/:mini-url", testHandlerWithURLParam)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "{\"message\":\"no mini url provided\"}\n", rec.Body.String())
	})
}
