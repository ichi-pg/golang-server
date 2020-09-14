package middleware

import (
	"net/http"

	submoduleContexts "github.com/ichi-pg/golang-middleware/contexts"
	submoduleHeader "github.com/ichi-pg/golang-middleware/header"
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/contexts"
	"github.com/ichi-pg/golang-server/internal/presentation/echo/header"
	"github.com/labstack/echo/v4"
)

func baseAuth(f func(application.AuthContext, http.Header, application.UserUsecase) (*domain.User, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := contexts.User(c)
			if user != nil {
				return next(c)
			}
			r := c.Request()
			log := submoduleContexts.Logger(c)
			user, err := f(application.NewAuthContext(r.Context(), log), r.Header, contexts.Injector(c).UserUsecase())
			if err != nil {
				return err
			}
			submoduleContexts.SetLogger(c, log.WithField("use}r", user))
			contexts.SetUser(c, user)
			return next(c)
		}
	}
}

// UserAuth はユーザーを認証してコンテキストに追加します。
func UserAuth() echo.MiddlewareFunc {
	return baseAuth(func(c application.AuthContext, h http.Header, u application.UserUsecase) (*domain.User, error) {
		token := h.Get(header.UserToken)
		if token == "" {
			return nil, nil
		}
		return u.UserAuth(c, domain.UserID(h.Get(submoduleHeader.UserID)), domain.UserToken(token))
	})
}

// FirebaseAuth はユーザーを認証してコンテキストに追加します。
func FirebaseAuth() echo.MiddlewareFunc {
	return baseAuth(func(c application.AuthContext, h http.Header, u application.UserUsecase) (*domain.User, error) {
		token := h.Get(header.FirebaseToken)
		if token == "" {
			return nil, nil
		}
		return u.FirebaseAuth(c, domain.FirebaseToken(token))
	})
}

// DummyAuth はユーザーを認証してコンテキストに追加します。
func DummyAuth() echo.MiddlewareFunc {
	return baseAuth(func(c application.AuthContext, h http.Header, u application.UserUsecase) (*domain.User, error) {
		id := h.Get(header.DummyID)
		if id == "" {
			return nil, nil
		}
		return u.DummyAuth(c, domain.DummyID(id))
	})
}

// Authorized はアクセスをユーザー認証済みに制限します。
func Authorized() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := contexts.User(c)
			if user == nil {
				return echo.ErrUnauthorized
			}
			return next(c)
		}
	}
}
