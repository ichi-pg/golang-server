package resolver

import (
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/domain"
	"github.com/sirupsen/logrus"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver はHTTPリスナーとの依存関係を解消します。
type Resolver struct {
	Logger   *logrus.Entry
	User     *domain.User
	Injector application.Injector
}

// NewResolver はResolverの生成方法を制限します。
func NewResolver(log *logrus.Entry, user *domain.User, i application.Injector) *Resolver {
	return &Resolver{
		Logger:   log,
		User:     user,
		Injector: i,
	}
}
