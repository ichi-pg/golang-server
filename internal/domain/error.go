package domain

import (
	"net/http"
)

var (
	// ErrNoSuchEntity is returned when no entity was found for a given key.
	ErrNoSuchEntity = NewRequestError(http.StatusNotFound, "404 no such entity")
)

// RequestError は外部要因エラーの識別に使用します。
type RequestError struct {
	Code int
	Msg  string
}

// NewRequestError は外部要因エラーを生成します。
func NewRequestError(code int, msg ...string) error {
	err := RequestError{
		Code: code,
	}
	if len(msg) > 0 {
		err.Msg = msg[0]
	}
	return err
}

func (e RequestError) Error() string {
	return e.Msg
}
