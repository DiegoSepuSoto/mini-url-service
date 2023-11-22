package miniurl

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *miniURLHandler) ServeMiniURLHandler(c echo.Context) error {
	ctx := context.Background()
	miniURL := c.Param("mini-url")

	minifiedURLResponse, err := h.miniURLUseCase.GetMinifiedURL(ctx, miniURL)
	if err != nil {
		log.Errorf("sending user to default url, error: %s", err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, os.Getenv("DEFAULT_URL_REDIRECT"))
	}

	log.WithFields(
		log.Fields{"originalURL": minifiedURLResponse.MinifiedURL, "miniURL": miniURL},
	).Info("mini url served successfully")

	return c.Redirect(http.StatusMovedPermanently, minifiedURLResponse.MinifiedURL)
}
