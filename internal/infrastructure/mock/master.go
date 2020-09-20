package mock

import (
	"context"

	"github.com/ichi-pg/golang-server/internal/domain"
)

type masterRepository struct {
}

// MasterRepository はMasterRepositoryのモック実装を返します。
func MasterRepository() domain.MasterRepository {
	return masterRepository{}
}

func (r masterRepository) PaymentItems(c context.Context) ([]domain.PaymentItem, error) {
	panic("TODO")
}

func (r masterRepository) PaymentItem(c context.Context, id domain.PaymentItemID) (*domain.PaymentItem, error) {
	panic("TODO")
}
