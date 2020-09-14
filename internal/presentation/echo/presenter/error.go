package presenter

import (
	"github.com/ichi-pg/golang-middleware/presenter"
	"github.com/labstack/echo/v4"
)

type errorPresenter struct {
}

// ErrorPresenter はエラーレスポンスを実装します。
func ErrorPresenter() presenter.ErrorPresenter {
	return errorPresenter{}
}

func (ep errorPresenter) Response(c echo.Context, err error) error {
	// TODO
	// code := http.StatusInternalServerError
	// msg := err.Error()
	// if reqErr, ok := err.(domain.RequestError); ok {
	// 	code = reqErr.Code
	// }
	// if httpErr, ok := err.(*echo.HTTPError); ok {
	// 	code = httpErr.Code
	// 	msg = fmt.Sprintf("%v", httpErr.Message)
	// }
	// c.JSON(code,
	// 	restError{
	// 		Errors: []restErrorInfo{
	// 			{
	// 				Message: msg,
	// 				Extensions: restErrorExtensions{
	// 					Status: code,
	// 				},
	// 			},
	// 		},
	// 	},
	// )
	return err
}
