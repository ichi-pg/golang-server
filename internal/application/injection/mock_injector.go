package injection

import (
	"github.com/ichi-pg/golang-server/internal/application"
	"github.com/ichi-pg/golang-server/internal/infrastructure/mock"
)

type mockInjector struct {
}

// MockInjector はapplication層以下のモックを注入します。
func MockInjector() application.Injector {
	return mockInjector{}
}

func (i mockInjector) UserUsecase() application.UserUsecase {
	return application.NewUserUsecase(
		mock.UserRepository(),
		mock.FirebaseRepository(),
	)
}

func (i mockInjector) RankingUsecase() application.RankingUsecase {
	return application.NewRankingUsecase(
		mock.RankingRepository(),
	)
}

func (i mockInjector) PaymentUsecase() application.PaymentUsecase {
	return application.NewPaymentUsecase(
		mock.PaymentRepository(),
	)
}
