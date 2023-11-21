package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func URLParamMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		miniURL := c.Param("mini-url")

		if miniURL == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": "no mini url provided"})
		}

		return next(c)
	}
}
