package miniurl

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/DiegoSepuSoto/mini-url-service/src/shared"
)

// GetMinifiedURL godoc
// @Summary      Get Minified URL
// @Description  Returns as an API Response the stored minified URL from mini URL provided
// @Tags         MiniURL
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token"
// @Success      200  {object}  models.MinifiedURLResponse "OK"
// @Failure      401  {object}  shared.EchoErrorResponse "Unauthorized"
// @Failure      500  {object}  shared.EchoErrorResponse "Application Error"
// @Router       /api/{mini-url} [get]
func (h *miniURLHandler) GetMinifiedURL(c echo.Context) error {
	ctx := context.Background()
	miniURL := c.Param("mini-url")

	minifiedURLResponse, err := h.miniURLUseCase.GetMinifiedURL(ctx, miniURL)
	if err != nil {
		log.Error(err)
		return c.JSON(shared.GetHTTPStatusErrorCode(err),
			&shared.EchoErrorResponse{Message: "error retrieving minified url"})
	}

	log.WithFields(
		log.Fields{"originalURL": minifiedURLResponse.MinifiedURL, "miniURL": miniURL},
	).Info("mini url sent successfully")

	return c.JSON(http.StatusOK, minifiedURLResponse)
}
