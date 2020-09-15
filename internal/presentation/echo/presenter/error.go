package presenter

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ichi-pg/golang-middleware/presenter"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type errorPresenter struct {
}

// ErrorPresenter はエラーレスポンスを実装します。
func ErrorPresenter() presenter.ErrorPresenter {
	return errorPresenter{}
}

func (ep errorPresenter) Response(c echo.Context, err error) error {
	code := http.StatusInternalServerError
	msg := err.Error()
	if reqErr, ok := err.(domain.RequestError); ok {
		code = reqErr.Code
	}
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
		msg = fmt.Sprintf("%v", httpErr.Message)
	}
	c.JSON(code,
		graphql.Response{
			Errors: gqlerror.List{
				{
					Message: msg,
					Extensions: map[string]interface{}{
						"status": code,
					},
				},
			},
		},
	)
	return err
}
