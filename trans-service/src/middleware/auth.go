package middleware

import (
	"github.com/labstack/echo/v4"
)

const (
	HEADER_KEY_TOKEN = "authToken"
)

func AuthChecker() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 从header中获取token
			token := c.Request().Header.Get(HEADER_KEY_TOKEN)
			if token == "" {
				return echo.ErrUnauthorized
			}

			// TODO check token is valid
			return next(c)
		}
	}
}
