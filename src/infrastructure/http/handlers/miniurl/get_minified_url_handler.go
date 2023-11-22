package miniurl

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (h *miniURLHandler) GetMinifiedURL(c echo.Context) error {
	ctx := context.Background()
	miniURL := c.Param("mini-url")

	minifiedURLResponse, err := h.miniURLUseCase.GetMinifiedURL(ctx, miniURL)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "error retrieving minified url"})
	}

	return c.JSON(http.StatusOK, minifiedURLResponse)
}
