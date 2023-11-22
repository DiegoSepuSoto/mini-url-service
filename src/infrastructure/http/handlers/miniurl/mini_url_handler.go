package miniurl

import (
	"github.com/labstack/echo/v4"

	"github.com/DiegoSepuSoto/mini-url-service/src/application/usecase"
	"github.com/DiegoSepuSoto/mini-url-service/src/infrastructure/http/handlers/middlewares"
)

type miniURLHandler struct {
	miniURLUseCase usecase.MiniURLUseCase
}

func NewMiniURLHandler(e *echo.Echo, miniURLUseCase usecase.MiniURLUseCase) *miniURLHandler {
	h := &miniURLHandler{
		miniURLUseCase: miniURLUseCase,
	}

	g := e.Group("/api", middlewares.APIMetricsMiddleware, middlewares.AuthMiddleware, middlewares.URLParamMiddleware)
	g.GET("/:mini-url", h.GetMinifiedURL)

	f := e.Group("", middlewares.RedirectMetricsMiddleware, middlewares.URLParamMiddleware)
	f.GET("/:mini-url", h.ServeMiniURLHandler)

	return h
}
