package application

// Injector はapplication層以下の依存関係を抽象化します。
type Injector interface {
	UserUsecase() UserUsecase
	RankingUsecase() RankingUsecase
	PaymentUsecase() PaymentUsecase
}
