package datastore

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
)

type paymentRepository struct {
}

// PaymentRepository はPaymentRepositoryのDatastore実装を返します。
func PaymentRepository() domain.PaymentRepository {
	return paymentRepository{}
}

func (r paymentRepository) Pay(c context.Context, user *domain.User, paymentItemID domain.PaymentItemID) (*domain.PaymentLog, error) {
	panic("TODO")
}

func (r paymentRepository) Logs(c context.Context, user *domain.User, cursor domain.Cursor) ([]domain.PaymentLog, domain.Cursor, error) {
	panic("TODO")
}

//TODO Test code
