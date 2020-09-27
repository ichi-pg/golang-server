package mock

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
)

type paymentRepository struct {
}

// PaymentRepository はPaymentRepositoryのモック実装を返します。
func PaymentRepository() domain.PaymentRepository {
	return paymentRepository{}
}

func (r paymentRepository) Pay(c context.Context, user *domain.User, paymentItem *domain.PaymentItem) (*domain.PaymentLog, error) {
	panic("TODO")
}

func (r paymentRepository) Logs(c context.Context, user *domain.User, cursor domain.Cursor) ([]domain.PaymentLog, domain.Cursor, error) {
	panic("TODO")
}
