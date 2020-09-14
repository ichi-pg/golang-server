package application

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/sirupsen/logrus"
)

// AuthContext は認証ユースケースのためのコンテキストを伝播します。
type AuthContext struct {
	Context context.Context
	Logger  *logrus.Entry
}

// UserContext はユーザーユースケースのためのコンテキストを伝播します。
type UserContext struct {
	Context context.Context
	Logger  *logrus.Entry
	User    *domain.User
}

// NewAuthContext は認証ユースケースのためのコンテキストを生成します。
func NewAuthContext(c context.Context, log *logrus.Entry) AuthContext {
	return AuthContext{
		Context: c,
		Logger:  log,
	}
}

// NewUserContext はユーザーユースケースのためのコンテキストを生成します。
func NewUserContext(c context.Context, log *logrus.Entry, user *domain.User) UserContext {
	return UserContext{
		Context: c,
		Logger:  log,
		User:    user,
	}
}
