package datastore

import (
	"github.com/ichi-pg/golang-server/internal/domain"
)

type paymentRepository struct {
}

// PaymentRepository はPaymentRepositoryのDatastore実装を返します。
func PaymentRepository() domain.PaymentRepository {
	return paymentRepository{}
}

//TODO Test code
