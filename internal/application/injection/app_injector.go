package injection

import (
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/infrastructure/datastore"
	"github.com/ichi-pg/golang-server/internal/infrastructure/firebases"
)

type appInjector struct {
}

// AppInjector はapplication層以下の依存関係を注入します。
func AppInjector() application.Injector {
	return appInjector{}
}

func (i appInjector) UserUsecase() application.UserUsecase {
	return application.NewUserUsecase(
		datastore.UserRepository(),
		firebases.FirebaseRepository(),
	)
}
