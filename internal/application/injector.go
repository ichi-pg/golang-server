package application

// Injector はapplication層以下の依存関係を抽象化します。
type Injector interface {
	UserUsecase() UserUsecase
	RankingUsecase() RankingUsecase
}

//TODO 購入モック（マスターデータ、アイテムリスト、Paymentログ、購入履歴ページング）
