package middleware

import (
	"github.com/ichi-pg/golang-server/internal/application/injection"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/contexts"
	"github.com/labstack/echo/v4"
)

// Injector は依存関係をコンテキストに追加します。
func Injector() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			contexts.SetInjector(c, injection.AppInjector())
			return next(c)
		}
	}
}
