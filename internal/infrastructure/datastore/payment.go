package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/ichi-pg/golang-server/internal/domain"
)

const paymentLogKind = "User"

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
	logs := []domain.PaymentLog{}
	nextCursor, err := runPagination(c,
		newQuery(paymentLogKind).Ancestor(newKey(userKind, string(user.ID), nil)),
		20, cursor,
		func(it *datastore.Iterator) error {
			log := domain.PaymentLog{}
			_, err := it.Next(&log)
			logs = append(logs, log)
			return err
		})
	return logs, nextCursor, err
}
