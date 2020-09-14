package contexts

import (
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/labstack/echo/v4"
)

const (
	user     = "user"
	injector = "injector"
)

// SetUser はユーザーをコンテキストに追加します。
func SetUser(c echo.Context, u *domain.User) {
	c.Set(user, u)
}

// User はユーザーをコンテキストから取り出します。
func User(c echo.Context) *domain.User {
	return c.Get(user).(*domain.User)
}

// SetInjector は依存関係をコンテキストに追加します。
func SetInjector(c echo.Context, i application.Injector) {
	c.Set(injector, i)
}

// Injector は依存関係をコンテキストから取り出します。
func Injector(c echo.Context) application.Injector {
	return c.Get(injector).(application.Injector)
}
