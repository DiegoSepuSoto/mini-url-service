package miniurl

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// ServeMiniURLHandler godoc
// @Summary      Serve Minified URL
// @Description  Serves on the browser the stored minified URL from mini URL provided
// @Tags         MiniURL
// @Success      301  {object}  models.MinifiedURLResponse "Full Redirect"
// @Failure      307  {object}  shared.EchoErrorResponse "Temporary Redirect"
// @Router       /{mini-url} [get]
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
