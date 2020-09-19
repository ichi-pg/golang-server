package mock

import (
	"github.com/ichi-pg/golang-server/internal/domain"
)

type paymentRepository struct {
}

// PaymentRepository はPaymentRepositoryのモック実装を返します。
func PaymentRepository() domain.PaymentRepository {
	return paymentRepository{}
}
