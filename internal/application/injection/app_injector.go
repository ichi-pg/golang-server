package injection

import (
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/infrastructure/datastore"
	"github.com/ichi-pg/golang-server/internal/infrastructure/firebases"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
	"github.com/ichi-pg/golang-server/internal/infrastructure/redis"
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

func (i appInjector) RankingUsecase() application.RankingUsecase {
	return application.NewRankingUsecase(
		redis.RankingRepository(datastore.UserRepository()),
	)
}

func (i appInjector) PaymentUsecase() application.PaymentUsecase {
	return application.NewPaymentUsecase(
		datastore.PaymentRepository(mock.MasterRepository()),
		mock.MasterRepository(),
	)
}
