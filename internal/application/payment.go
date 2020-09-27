package application

import (
	"github.com/ichi-pg/golang-server/internal/domain"
)

// PaymentUsecase は課金のユースケースを実装します。
type PaymentUsecase struct {
	repo       domain.PaymentRepository
	masterRepo domain.MasterRepository
}

// NewPaymentUsecase は課金のユースケースを生成します。
func NewPaymentUsecase(repo domain.PaymentRepository, masterRepo domain.MasterRepository) PaymentUsecase {
	return PaymentUsecase{
		repo:       repo,
		masterRepo: masterRepo,
	}
}

// Items は品揃えを取得します。
func (u PaymentUsecase) Items(c AuthContext) ([]domain.PaymentItem, error) {
	return u.masterRepo.PaymentItems(c.Context)
}

// Pay はアイテムを購入します。
func (u PaymentUsecase) Pay(c UserContext, id domain.PaymentItemID) (*domain.PaymentLog, error) {
	paymentItem, err := u.masterRepo.PaymentItem(c.Context, id)
	if err != nil {
		return nil, err
	}
	return u.repo.Pay(c.Context, c.User, paymentItem)
}

// Logs は購入履歴を取得します。
func (u PaymentUsecase) Logs(c UserContext, cursor domain.Cursor) ([]domain.PaymentLog, domain.Cursor, error) {
	return u.repo.Logs(c.Context, c.User, cursor)
}

//TODO Test code
