package handler

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	submoduleContexts "github.com/ichi-pg/golang-middleware/contexts"
	"github.com/ichi-pg/golang-middleware/util"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/contexts"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/generated"
	"github.com/ichi-pg/golang-server/internal/presentation/graph/resolver"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQuery はGraphQLをリクエストします。
func GraphQuery(c echo.Context) error {
	log := submoduleContexts.Logger(c)
	user := contexts.User(c)
	i := contexts.Injector(c)
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver.NewResolver(log, user, i)},
		),
	)
	w := util.MaxStatusResponseWriter(c.Response())
	server.SetErrorPresenter(
		func(ctx context.Context, e error) *gqlerror.Error {
			if reqErr, ok := e.(domain.RequestError); ok {
				w.WriteHeader(reqErr.Code)
				return &gqlerror.Error{
					Message: e.Error(),
					Path:    graphql.GetFieldContext(ctx).Path(),
					Extensions: map[string]interface{}{
						"status": reqErr.Code,
					},
				}
			}
			err := graphql.DefaultErrorPresenter(ctx, e)
			if _, ok := e.(*gqlerror.Error); ok {
				w.WriteHeader(http.StatusBadRequest)
				return err
			}
			if _, ok := e.(graphql.ExtendedError); ok {
				w.WriteHeader(http.StatusBadRequest)
				return err
			}
			w.WriteHeader(http.StatusInternalServerError)
			return err
		},
	)
	server.ServeHTTP(w, c.Request())
	return nil
}
