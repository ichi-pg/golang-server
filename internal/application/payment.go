package application

import (
	"github.com/ichi-pg/golang-server/internal/domain"
)

// PaymentUsecase は課金のユースケースを実装します。
type PaymentUsecase struct {
	repo domain.PaymentRepository
}

// NewPaymentUsecase は課金のユースケースを生成します。
func NewPaymentUsecase(repo domain.PaymentRepository) PaymentUsecase {
	return PaymentUsecase{
		repo: repo,
	}
}

// Items は品揃えを取得します。
func (u PaymentUsecase) Items(c AuthContext) error {
	//TODO
	return nil
}

// Pay はアイテムを購入します。
func (u PaymentUsecase) Pay(c UserContext) error {
	//TODO KPIログ
	return nil
}

// Logs は購入履歴を取得します。
func (u PaymentUsecase) Logs(c UserContext) error {
	//TODO
	return nil
}

//TODO Test code
