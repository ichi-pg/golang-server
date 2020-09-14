package echo

import (
	"os"

	submoduleMiddlewre "github.com/ichi-pg/golang-middleware/middleware"
	"github.com/ichi-pg/golang-server/internal/pkg/env"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/handler"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/middleware"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/presenter"
	"github.com/labstack/echo/v4"
)

// Start はechoサーバーを起動します。
func Start() error {
	e := echo.New()
	e.HideBanner = true

	ep := presenter.ErrorPresenter()

	e.Use(submoduleMiddlewre.Recover(ep))
	e.Use(submoduleMiddlewre.Logger(ep))
	e.Use(middleware.Injector())

	{
		graph := e.Group("/api/graph")
		//TODO
		// graph.Use(submoduleMiddlewre.Maintenance())
		// graph.Use(submoduleMiddlewre.ClientVersion())
		graph.Use(middleware.DummyAuth(), middleware.FirebaseAuth(), middleware.UserAuth())
		graph.Use(middleware.Authorized())
		graph.POST("/query", handler.GraphQuery)
	}

	return e.Start(":" + os.Getenv(env.Port))
}
